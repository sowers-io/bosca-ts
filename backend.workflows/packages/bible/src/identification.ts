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

export type BookIdentificationCode = 'GEN' // Genesis
  | 'EXO' // Exodus
  | 'LEV' // Leviticus
  | 'NUM' // Numbers
  | 'DEU' // Deuteronomy
  | 'JOS' // Joshua
  | 'JDG' // Judges
  | 'RUT' // Ruth
  | '1SA' // 1 Samuel
  | '2SA' // 2 Samuel
  | '1KI' // 1 Kings
  | '2KI' // 2 Kings
  | '1CH' // 1 Chronicles
  | '2CH' // 2 Chronicles
  | 'EZR' // Ezra
  | 'NEH' // Nehemiah
  | 'EST' // Esther (Hebrew)
  | 'JOB' // Job
  | 'PSA' // Psalms
  | 'PRO' // Proverbs
  | 'ECC' // Ecclesiastes
  | 'SNG' // Song of Songs
  | 'ISA' // Isaiah
  | 'JER' // Jeremiah
  | 'LAM' // Lamentations
  | 'EZK' // Ezekiel
  | 'DAN' // Daniel (Hebrew)
  | 'HOS' // Hosea
  | 'JOL' // Joel
  | 'AMO' // Amos
  | 'OBA' // Obadiah
  | 'JON' // Jonah
  | 'MIC' // Micah
  | 'NAM' // Nahum
  | 'HAB' // Habakkuk
  | 'ZEP' // Zephaniah
  | 'HAG' // Haggai
  | 'ZEC' // Zechariah
  | 'MAL' // Malachi
  | 'MAT' // Matthew
  | 'MRK' // Mark
  | 'LUK' // Luke
  | 'JHN' // John
  | 'ACT' // Acts
  | 'ROM' // Romans
  | '1CO' // 1 Corinthians
  | '2CO' // 2 Corinthians
  | 'GAL' // Galatians
  | 'EPH' // Ephesians
  | 'PHP' // Philippians
  | 'COL' // Colossians
  | '1TH' // 1 Thessalonians
  | '2TH' // 2 Thessalonians
  | '1TI' // 1 Timothy
  | '2TI' // 2 Timothy
  | 'TIT' // Titus
  | 'PHM' // Philemon
  | 'HEB' // Hebrews
  | 'JAS' // James
  | '1PE' // 1 Peter
  | '2PE' // 2 Peter
  | '1JN' // 1 John
  | '2JN' // 2 John
  | '3JN' // 3 John
  | 'JUD' // Jude
  | 'REV' // Revelation
  | 'TOB' // Tobit
  | 'JDT' // Judith
  | 'ESG' // Esther Greek
  | 'WIS' // Wisdom of Solomon
  | 'SIR' // Sirach (Ecclesiasticus)
  | 'BAR' // Baruch
  | 'LJE' // Letter of Jeremiah
  | 'S3Y' // Song of 3 Young Men
  | 'SUS' // Susanna
  | 'BEL' // Bel and the Dragon
  | '1MA' // 1 Maccabees
  | '2MA' // 2 Maccabees
  | '3MA' // 3 Maccabees
  | '4MA' // 4 Maccabees
  | '1ES' // 1 Esdras (Greek)
  | '2ES' // 2 Esdras (Latin)
  | 'MAN' // Prayer of Manasseh
  | 'PS2' // Psalm 151
  | 'ODA' // Odes
  | 'PSS' // Psalms of Solomon
  | 'EZA' // Apocalypse of Ezra
  | '5EZ' // 5 Ezra
  | '6EZ' // 6 Ezra
  | 'DAG' // Daniel Greek
  | 'PS3' // Psalms 152-155
  | '2BA' // 2 Baruch (Apocalypse)
  | 'LBA' // Letter of Baruch
  | 'JUB' // Jubilees
  | 'ENO' // Enoch
  | '1MQ' // 1 Meqabyan
  | '2MQ' // 2 Meqabyan
  | '3MQ' // 3 Meqabyan
  | 'REP' // Reproof
  | '4BA' // 4 Baruch
  | 'LAO' // Laodiceans

export const BookIdentificationCodes: BookIdentificationCode[] = ['GEN', 'EXO', 'LEV', 'NUM', 'DEU', 'JOS', 'JDG', 'RUT', '1SA', '2SA', '1KI', '2KI', '1CH', '2CH', 'EZR', 'NEH', 'EST', 'JOB', 'PSA', 'PRO', 'ECC', 'SNG', 'ISA', 'JER', 'LAM', 'EZK', 'DAN', 'HOS', 'JOL', 'AMO', 'OBA', 'JON', 'MIC', 'NAM', 'HAB', 'ZEP', 'HAG', 'ZEC', 'MAL', 'MAT', 'MRK', 'LUK', 'JHN', 'ACT', 'ROM', '1CO', '2CO', 'GAL', 'EPH', 'PHP', 'COL', '1TH', '2TH', '1TI', '2TI', 'TIT', 'PHM', 'HEB', 'JAS', '1PE', '2PE', '1JN', '2JN', '3JN', 'JUD', 'REV', 'TOB', 'JDT', 'ESG', 'WIS', 'SIR', 'BAR', 'LJE', 'S3Y', 'SUS', 'BEL', '1MA', '2MA', '3MA', '4MA', '1ES', '2ES', 'MAN', 'PS2', 'ODA', 'PSS', 'EZA', '5EZ', '6EZ', 'DAG', 'PS3', '2BA', 'LBA', 'JUB', 'ENO', '1MQ', '2MQ', '3MQ', 'REP', '4BA', 'LAO']