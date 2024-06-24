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

package usx

import "encoding/xml"

type BookCode int

const (
	BookUnknown BookCode = iota
	// BookGEN Genesis
	BookGEN
	// BookEXO Exodus
	BookEXO
	// BookLEV Leviticus
	BookLEV
	// BookNUM Numbers
	BookNUM
	// BookDEU Deuteronomy
	BookDEU
	// BookJOS Joshua
	BookJOS
	// BookJDG Judges
	BookJDG
	// BookRUT Ruth
	BookRUT
	// Book1SA 1 Samuel
	Book1SA
	// Book2SA 2 Samuel
	Book2SA
	// Book1KI 1 Kings
	Book1KI
	// Book2KI 2 Kings
	Book2KI
	// Book1CH 1 Chronicles
	Book1CH
	// Book2CH 2 Chronicles
	Book2CH
	// BookEZR Ezra
	BookEZR
	// BookNEH Nehemiah
	BookNEH
	// BookEST Esther
	BookEST
	// BookJOB Job
	BookJOB
	// BookPSA Psalms
	BookPSA
	// BookPRO Proverbs
	BookPRO
	// BookECC Ecclesiastes
	BookECC
	// BookSNG Song of Songs
	BookSNG
	// BookISA Isaiah
	BookISA
	// BookJER Jeremiah
	BookJER
	// BookLAM Lamentations
	BookLAM
	// BookEZK Ezekiel
	BookEZK
	// BookDAN Daniel (Hebrew)
	BookDAN
	// BookHOS Hosea
	BookHOS
	// BookJOL Joel
	BookJOL
	// BookAMO Amos
	BookAMO
	// BookOBA Obadiah
	BookOBA
	// BookJON Jonah
	BookJON
	// BookMIC Micah
	BookMIC
	// BookNAM Nahum
	BookNAM
	// BookHAB Habakkuk
	BookHAB
	// BookZEP Zephaniah
	BookZEP
	// BookHAG Haggai
	BookHAG
	// BookZEC Zechariah
	BookZEC
	// BookMAL Malachi
	BookMAL
	// BookMAT Matthew
	BookMAT
	// BookMRK Mark
	BookMRK
	// BookLUK Luke
	BookLUK
	// BookJHN John
	BookJHN
	// BookACT Acts
	BookACT
	// BookROM Romans
	BookROM
	// Book1CO 1 Corinthians
	Book1CO
	// Book2CO 2 Corinthians
	Book2CO
	// BookGAL Galatians
	BookGAL
	// BookEPH Ephesians
	BookEPH
	// BookPHP Philippians
	BookPHP
	// BookCOL Colossians
	BookCOL
	// Book1TH 1 Thessalonians
	Book1TH
	// Book2TH 2 Thessalonians
	Book2TH
	// Book1TI 1 Timothy
	Book1TI
	// Book2TI 2 Timothy
	Book2TI
	// BookTIT Titus
	BookTIT
	// BookPHM Philemon
	BookPHM
	// BookHEB Hebrews
	BookHEB
	// BookJAS James
	BookJAS
	// Book1PE 1 Peter
	Book1PE
	// Book2PE 2 Peter
	Book2PE
	// Book1JN 1 John
	Book1JN
	// Book2JN 2 John
	Book2JN
	// Book3JN 3 John
	Book3JN
	// BookJUD Jude
	BookJUD
	// BookREV Revelation
	BookREV
	// BookTOB Tobit
	BookTOB
	// BookJDT Judith
	BookJDT
	// BookESG Esther Greek
	BookESG
	// BookWIS Wisdom of Solomon
	BookWIS
	// BookSIR Sirach (Ecclesiasticus)
	BookSIR
	// BookBAR Baruch
	BookBAR
	// BookLJE Letter of Jeremiah
	BookLJE
	// BookS3Y Song of 3 Young Men
	BookS3Y
	// BookSUS Susanna
	BookSUS
	// BookBEL Bel and the Dragon
	BookBEL
	// Book1MA 1 Maccabees
	Book1MA
	// Book2MA 2 Maccabees
	Book2MA
	// Book3MA 3 Maccabees
	Book3MA
	// Book4MA 4 Maccabees
	Book4MA
	// Book1ES 1 Esdras (Greek)
	Book1ES
	// Book2ES 2 Esdras (Latin)
	Book2ES
	// BookMAN Prayer of Manasseh
	BookMAN
	// BookPS2 Psalm 151
	BookPS2
	// BookODA Odes
	BookODA
	// BookPSS Psalms of Solomon
	BookPSS
	// BookEZA Apocalypse of Ezra
	BookEZA
	// Book5EZ 5 Ezra
	Book5EZ
	// Book6EZ 6 Ezra
	Book6EZ
	// BookDAG Daniel Greek
	BookDAG
	// BookPS3 Psalms 152-155
	BookPS3
	// Book2BA 2 Baruch (Apocalypse)
	Book2BA
	// BookLBA Letter of Baruch
	BookLBA
	// BookJUB Jubilees
	BookJUB
	// BookENO Enoch
	BookENO
	// Book1MQ 1 Meqabyan
	Book1MQ
	// Book2MQ 2 Meqabyan
	Book2MQ
	// Book3MQ 3 Meqabyan
	Book3MQ
	// BookREP Reproof
	BookREP
	// Book4BA 4 Baruch
	Book4BA
	// BookLAO LAO
	BookLAO
)

func (c BookCode) ToString() string {
	switch c {
	case BookGEN:
		return "GEN"
	case BookEXO:
		return "EXO"
	case BookLEV:
		return "LEV"
	case BookNUM:
		return "NUM"
	case BookDEU:
		return "DEU"
	case BookJOS:
		return "JOS"
	case BookJDG:
		return "JDG"
	case BookRUT:
		return "RUT"
	case Book1SA:
		return "1SA"
	case Book2SA:
		return "2SA"
	case Book1KI:
		return "1KI"
	case Book2KI:
		return "2KI"
	case Book1CH:
		return "1CH"
	case Book2CH:
		return "2CH"
	case BookEZR:
		return "EZR"
	case BookNEH:
		return "NEH"
	case BookEST:
		return "EST"
	case BookJOB:
		return "JOB"
	case BookPSA:
		return "PSA"
	case BookPRO:
		return "PRO"
	case BookECC:
		return "ECC"
	case BookSNG:
		return "SNG"
	case BookISA:
		return "ISA"
	case BookJER:
		return "JER"
	case BookLAM:
		return "LAM"
	case BookEZK:
		return "EZK"
	case BookDAN:
		return "DAN"
	case BookHOS:
		return "HOS"
	case BookJOL:
		return "JOL"
	case BookAMO:
		return "AMO"
	case BookOBA:
		return "OBA"
	case BookJON:
		return "JON"
	case BookMIC:
		return "MIC"
	case BookNAM:
		return "NAM"
	case BookHAB:
		return "HAB"
	case BookZEP:
		return "ZEP"
	case BookHAG:
		return "HAG"
	case BookZEC:
		return "ZEC"
	case BookMAL:
		return "MAL"
	case BookMAT:
		return "MAT"
	case BookMRK:
		return "MRK"
	case BookLUK:
		return "LUK"
	case BookJHN:
		return "JHN"
	case BookACT:
		return "ACT"
	case BookROM:
		return "ROM"
	case Book1CO:
		return "1CO"
	case Book2CO:
		return "2CO"
	case BookGAL:
		return "GAL"
	case BookEPH:
		return "EPH"
	case BookPHP:
		return "PHP"
	case BookCOL:
		return "COL"
	case Book1TH:
		return "1TH"
	case Book2TH:
		return "2TH"
	case Book1TI:
		return "1TI"
	case Book2TI:
		return "2TI"
	case BookTIT:
		return "TIT"
	case BookPHM:
		return "PHM"
	case BookHEB:
		return "HEB"
	case BookJAS:
		return "JAS"
	case Book1PE:
		return "1PE"
	case Book2PE:
		return "2PE"
	case Book1JN:
		return "1JN"
	case Book2JN:
		return "2JN"
	case Book3JN:
		return "3JN"
	case BookJUD:
		return "JUD"
	case BookREV:
		return "REV"
	case BookTOB:
		return "TOB"
	case BookJDT:
		return "JDT"
	case BookESG:
		return "ESG"
	case BookWIS:
		return "WIS"
	case BookSIR:
		return "SIR"
	case BookBAR:
		return "BAR"
	case BookLJE:
		return "LJE"
	case BookS3Y:
		return "S3Y"
	case BookSUS:
		return "SUS"
	case BookBEL:
		return "BEL"
	case Book1MA:
		return "1MA"
	case Book2MA:
		return "2MA"
	case Book3MA:
		return "3MA"
	case Book4MA:
		return "4MA"
	case Book1ES:
		return "1ES"
	case Book2ES:
		return "2ES"
	case BookMAN:
		return "MAN"
	case BookPS2:
		return "PS2"
	case BookODA:
		return "ODA"
	case BookPSS:
		return "PSS"
	case BookEZA:
		return "EZA"
	case Book5EZ:
		return "5EZ"
	case Book6EZ:
		return "6EZ"
	case BookDAG:
		return "DAG"
	case BookPS3:
		return "PS3"
	case Book2BA:
		return "2BA"
	case BookLBA:
		return "LBA"
	case BookJUB:
		return "JUB"
	case BookENO:
		return "ENO"
	case Book1MQ:
		return "1MQ"
	case Book2MQ:
		return "2MQ"
	case Book3MQ:
		return "3MQ"
	case BookREP:
		return "REP"
	case Book4BA:
		return "4BA"
	case BookLAO:
		return "LAO"
	default:
		return "UNKNOWN"
	}
}

func parseBookType(attr xml.Attr) BookCode {
	switch attr.Value {
	case "GEN":
		return BookGEN
	case "EXO":
		return BookEXO
	case "LEV":
		return BookLEV
	case "NUM":
		return BookNUM
	case "DEU":
		return BookDEU
	case "JOS":
		return BookJOS
	case "JDG":
		return BookJDG
	case "RUT":
		return BookRUT
	case "1SA":
		return Book1SA
	case "2SA":
		return Book2SA
	case "1KI":
		return Book1KI
	case "2KI":
		return Book2KI
	case "1CH":
		return Book1CH
	case "2CH":
		return Book2CH
	case "EZR":
		return BookEZR
	case "NEH":
		return BookNEH
	case "EST":
		return BookEST
	case "JOB":
		return BookJOB
	case "PSA":
		return BookPSA
	case "PRO":
		return BookPRO
	case "ECC":
		return BookECC
	case "SNG":
		return BookSNG
	case "ISA":
		return BookISA
	case "JER":
		return BookJER
	case "LAM":
		return BookLAM
	case "EZK":
		return BookEZK
	case "DAN":
		return BookDAN
	case "HOS":
		return BookHOS
	case "JOL":
		return BookJOL
	case "AMO":
		return BookAMO
	case "OBA":
		return BookOBA
	case "JON":
		return BookJON
	case "MIC":
		return BookMIC
	case "NAM":
		return BookNAM
	case "HAB":
		return BookHAB
	case "ZEP":
		return BookZEP
	case "HAG":
		return BookHAG
	case "ZEC":
		return BookZEC
	case "MAL":
		return BookMAL
	case "MAT":
		return BookMAT
	case "MRK":
		return BookMRK
	case "LUK":
		return BookLUK
	case "JHN":
		return BookJHN
	case "ACT":
		return BookACT
	case "ROM":
		return BookROM
	case "1CO":
		return Book1CO
	case "2CO":
		return Book2CO
	case "GAL":
		return BookGAL
	case "EPH":
		return BookEPH
	case "PHP":
		return BookPHP
	case "COL":
		return BookCOL
	case "1TH":
		return Book1TH
	case "2TH":
		return Book2TH
	case "1TI":
		return Book1TI
	case "2TI":
		return Book2TI
	case "TIT":
		return BookTIT
	case "PHM":
		return BookPHM
	case "HEB":
		return BookHEB
	case "JAS":
		return BookJAS
	case "1PE":
		return Book1PE
	case "2PE":
		return Book2PE
	case "1JN":
		return Book1JN
	case "2JN":
		return Book2JN
	case "3JN":
		return Book3JN
	case "JUD":
		return BookJUD
	case "REV":
		return BookREV
	case "TOB":
		return BookTOB
	case "JDT":
		return BookJDT
	case "ESG":
		return BookESG
	case "WIS":
		return BookWIS
	case "SIR":
		return BookSIR
	case "BAR":
		return BookBAR
	case "LJE":
		return BookLJE
	case "S3Y":
		return BookS3Y
	case "SUS":
		return BookSUS
	case "BEL":
		return BookBEL
	case "1MA":
		return Book1MA
	case "2MA":
		return Book2MA
	case "3MA":
		return Book3MA
	case "4MA":
		return Book4MA
	case "1ES":
		return Book1ES
	case "2ES":
		return Book2ES
	case "MAN":
		return BookMAN
	case "PS2":
		return BookPS2
	case "ODA":
		return BookODA
	case "PSS":
		return BookPSS
	case "EZA":
		return BookEZA
	case "5EZ":
		return Book5EZ
	case "6EZ":
		return Book6EZ
	case "DAG":
		return BookDAG
	case "PS3":
		return BookPS3
	case "2BA":
		return Book2BA
	case "LBA":
		return BookLBA
	case "JUB":
		return BookJUB
	case "ENO":
		return BookENO
	case "1MQ":
		return Book1MQ
	case "2MQ":
		return Book2MQ
	case "3MQ":
		return Book3MQ
	case "REP":
		return BookREP
	case "4BA":
		return Book4BA
	case "LAO":
		return BookLAO
	default:
		return BookUnknown
	}
}
