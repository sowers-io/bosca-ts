/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import ChildProcess, { ChildProcessWithoutNullStreams } from 'child_process'
import { logger } from '@bosca/common'
import process from 'process'

export type ProcessResult = { code: number, type: string, error?: string }

export async function transcribe(mp3File: string, language: string): Promise<any> {
  const spawned = ChildProcess.spawn(process.env.POETRY_PROGRAM!, ['run', 'python', 'main.py', process.cwd() + '/' + mp3File, language], {
    cwd: process.env.MEDIA_PYTHON_DIR!,
  })
  let json = ''
  let processingData = false
  await executeProcess(spawned, (data) => {
    if (processingData || data.startsWith('{')) {
      processingData = true
      json += data
    }
  })
  return JSON.parse(json)
}

export async function extractMP3(mp4File: string, mp3File: string) {
  const process = ChildProcess.spawn('ffmpeg', ['-i', mp4File, '-vn', mp3File])
  return await executeProcess(process, () => {})
}

export async function downloadHLS(url: string, outputFile: string): Promise<ProcessResult> {
  if (!URL.canParse(url)) throw new Error('invalid url')
  const process = ChildProcess.spawn('ffmpeg', ['-i', url, '-c', 'copy', outputFile])
  return await executeProcess(process, () => {})
}

async function executeProcess(process: ChildProcessWithoutNullStreams, onData: (data: string) => void): Promise<ProcessResult> {
  logger.debug({ pid: process.pid }, 'executing process')
  process.stdout.on('data', (data) => {
    onData(data.toString())
    logger.debug(data.toString())
  })
  process.stderr.on('data', (data) => {
    logger.debug(data.toString())
  })
  const result = await new Promise<ProcessResult>((resolve) => {
    let resolved = false
    process.on('error', (data) => {
      if (resolved) return
      resolved = true
      resolve({ 'code': -1, 'type': 'error', error: data.toString() } as ProcessResult)
    })
    process.on('disconnect', () => {
      if (resolved) return
      resolved = true
      resolve({ 'code': 0, 'type': 'disconnect' } as ProcessResult)
    })
    process.on('exit', (code) => {
      if (resolved) return
      resolved = true
      resolve({ 'code': code, 'type': 'exit' } as ProcessResult)
    })
    process.on('close', (code) => {
      if (resolved) return
      resolved = true
      resolve({ 'code': code, 'type': 'close' } as ProcessResult)
    })
  })
  process[Symbol.dispose]()
  return result
}