import whisper_timestamped as whisper
import json
import sys


def transcribe(file):
    audio = whisper.load_audio(file)
    model = whisper.load_model("tiny", device="cpu")
    result = whisper.transcribe(model, audio, language="en")
    print(json.dumps(result, indent = 0, ensure_ascii = False))
    sys.exit(0)


if __name__ == '__main__':
    transcribe(sys.argv[1])