import { test } from 'vitest';
import { downloadHLS, extractMP3, transcribe } from './ffmpeg';

test('test download mp4', async () => {
  await downloadHLS('https://devstreaming-cdn.apple.com/videos/streaming/examples/adv_dv_atmos/main.m3u8', 'test.mp4')
}, 120_000_000)

test('test convert to mp3', async () => {
  await extractMP3('test.mp4', 'test.mp3')
}, 120_000_000)

test('transcribe mp3', async () => {
  const data = await transcribe('test.mp3')
  console.log(data)
}, 120_000_000)