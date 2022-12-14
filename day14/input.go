package day14

const sample = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

const input0 = `496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
568,139 -> 573,139
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
534,78 -> 539,78
513,78 -> 518,78
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
553,133 -> 558,133
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
511,53 -> 515,53
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
543,144 -> 543,145 -> 560,145 -> 560,144
522,69 -> 527,69
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
560,133 -> 565,133
561,139 -> 566,139
518,94 -> 518,95 -> 532,95 -> 532,94
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
517,53 -> 521,53
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
518,94 -> 518,95 -> 532,95 -> 532,94
520,78 -> 525,78
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
546,114 -> 546,115 -> 554,115 -> 554,114
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
522,63 -> 527,63
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
564,136 -> 569,136
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
534,111 -> 534,112 -> 549,112 -> 549,111
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
526,72 -> 531,72
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
533,66 -> 538,66
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
527,78 -> 532,78
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
550,136 -> 555,136
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
556,130 -> 561,130
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
547,139 -> 552,139
505,57 -> 509,57
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
567,133 -> 572,133
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
520,55 -> 524,55
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
543,144 -> 543,145 -> 560,145 -> 560,144
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
575,139 -> 580,139
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
511,57 -> 515,57
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
514,55 -> 518,55
546,114 -> 546,115 -> 554,115 -> 554,114
557,136 -> 562,136
571,136 -> 576,136
546,114 -> 546,115 -> 554,115 -> 554,114
497,13 -> 497,17 -> 489,17 -> 489,22 -> 502,22 -> 502,17 -> 501,17 -> 501,13
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
525,60 -> 530,60
559,127 -> 564,127
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
517,81 -> 517,84 -> 509,84 -> 509,92 -> 524,92 -> 524,84 -> 522,84 -> 522,81
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
529,63 -> 534,63
526,66 -> 531,66
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
508,55 -> 512,55
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
534,111 -> 534,112 -> 549,112 -> 549,111
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
516,75 -> 521,75
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
523,75 -> 528,75
523,57 -> 527,57
519,66 -> 524,66
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
554,139 -> 559,139
543,144 -> 543,145 -> 560,145 -> 560,144
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
553,118 -> 553,120 -> 549,120 -> 549,124 -> 561,124 -> 561,120 -> 557,120 -> 557,118
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
514,51 -> 518,51
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
517,57 -> 521,57
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
518,94 -> 518,95 -> 532,95 -> 532,94
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
496,35 -> 496,29 -> 496,35 -> 498,35 -> 498,26 -> 498,35 -> 500,35 -> 500,25 -> 500,35 -> 502,35 -> 502,34 -> 502,35 -> 504,35 -> 504,31 -> 504,35 -> 506,35 -> 506,33 -> 506,35 -> 508,35 -> 508,25 -> 508,35
530,75 -> 535,75
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
519,72 -> 524,72
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
563,130 -> 568,130
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
531,98 -> 531,102 -> 526,102 -> 526,109 -> 540,109 -> 540,102 -> 535,102 -> 535,98
534,111 -> 534,112 -> 549,112 -> 549,111
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
561,171 -> 561,162 -> 561,171 -> 563,171 -> 563,168 -> 563,171 -> 565,171 -> 565,163 -> 565,171 -> 567,171 -> 567,163 -> 567,171 -> 569,171 -> 569,161 -> 569,171 -> 571,171 -> 571,170 -> 571,171 -> 573,171 -> 573,166 -> 573,171 -> 575,171 -> 575,161 -> 575,171 -> 577,171 -> 577,165 -> 577,171
501,48 -> 501,44 -> 501,48 -> 503,48 -> 503,39 -> 503,48 -> 505,48 -> 505,39 -> 505,48 -> 507,48 -> 507,41 -> 507,48 -> 509,48 -> 509,41 -> 509,48 -> 511,48 -> 511,47 -> 511,48 -> 513,48 -> 513,47 -> 513,48 -> 515,48 -> 515,43 -> 515,48
551,158 -> 551,154 -> 551,158 -> 553,158 -> 553,156 -> 553,158 -> 555,158 -> 555,152 -> 555,158 -> 557,158 -> 557,149 -> 557,158 -> 559,158 -> 559,156 -> 559,158 -> 561,158 -> 561,152 -> 561,158 -> 563,158 -> 563,155 -> 563,158 -> 565,158 -> 565,153 -> 565,158 -> 567,158 -> 567,155 -> 567,158 -> 569,158 -> 569,149 -> 569,158`
