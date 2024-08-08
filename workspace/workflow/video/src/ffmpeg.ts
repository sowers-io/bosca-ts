import ChildProcess, { ChildProcessWithoutNullStreams } from 'child_process';
import { logger } from '@bosca/common';
import process from 'process'

export type ProcessResult = { code: number, type: string, error?: string }

export async function transcribe(mp3File: string): Promise<any> {
  const spawned = ChildProcess.spawn('/home/brock/.local/bin/poetry', ['run', 'python', 'main.py', process.cwd() + '/' + mp3File], {
    cwd: '../video-py',
  })
  let json = ''
  await executeProcess(spawned, (data) => {
    if (data.startsWith('{')) {
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
    });
    process.on('close', (code) => {
      if (resolved) return
      resolved = true
      resolve({ 'code': code, 'type': 'close' } as ProcessResult)
    });
  })
  process[Symbol.dispose]()
  return result
}