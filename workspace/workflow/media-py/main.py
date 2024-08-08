#
# Copyright 2024 Sowers, LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

import whisper_timestamped as whisper
import json
import sys


def transcribe(file, language):
    audio = whisper.load_audio(file)
    model = whisper.load_model("tiny", device="cpu")
    result = whisper.transcribe(model, audio, language=language)
    print(json.dumps(result, indent = 0, ensure_ascii = False))
    sys.exit(0)


if __name__ == '__main__':
    transcribe(sys.argv[1], sys.argv[2])