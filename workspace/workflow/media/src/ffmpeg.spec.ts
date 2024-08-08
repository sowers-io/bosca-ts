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

import { test } from 'vitest'
import { downloadHLS, extractMP3, transcribe } from './ffmpeg'

test('test download mp4', async () => {
  await downloadHLS('https://devstreaming-cdn.apple.com/videos/streaming/examples/adv_dv_atmos/main.m3u8', 'test.mp4')
}, 120_000_000)

test('test convert to mp3', async () => {
  await extractMP3('test.mp4', 'test.mp3')
}, 120_000_000)

test('transcribe mp3', async () => {
  const data = await transcribe('test.mp3', 'my')
  console.log(data)
}, 120_000_000)