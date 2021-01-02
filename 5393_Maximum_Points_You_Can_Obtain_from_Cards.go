// 2020/4/26
package main

import (
	"fmt"
	// "reflect"
	"time"
)

func main() {
	start := time.Now()
	cardPoints, k := []int{1, 2, 3, 4, 5, 6, 1}, 3
	fmt.Println(maxScore(cardPoints, k) == 12)
	cardPoints, k = []int{2, 2, 2}, 2
	fmt.Println(maxScore(cardPoints, k) == 4)
	cardPoints, k = []int{9, 7, 7, 9, 7, 7, 9}, 7
	fmt.Println(maxScore(cardPoints, k) == 55)
	cardPoints, k = []int{1, 1000, 1}, 1
	fmt.Println(maxScore(cardPoints, k) == 1)
	cardPoints, k = []int{1, 79, 80, 1, 1, 1, 200, 1}, 3
	fmt.Println(maxScore(cardPoints, k) == 202)
	cardPoints, k = []int{1, 79, 80, 1, 1, 1, 200, 1}, 1
	fmt.Println(maxScore(cardPoints, k) == 1)
	cardPoints, k = []int{}, 0
	fmt.Println(maxScore(cardPoints, k) == 0)
	cardPoints, k = []int{100, 40, 17, 9, 73, 75}, 3
	fmt.Println(maxScore(cardPoints, k) == 248)
	cardPoints, k = []int{655, 282, 726, 868, 575, 957, 237, 228, 978, 391, 960, 968, 673, 165, 783, 287, 567, 427, 148, 662, 192, 713, 758, 772, 314, 15, 823, 160, 583, 364, 481, 474, 552, 756, 152, 547, 731, 713, 127, 779, 614, 941, 953, 102, 224, 243, 406, 855, 429, 227, 372, 298, 421, 982, 199, 546, 575, 897, 117, 506, 759, 838, 632, 842, 533, 505, 961, 689, 153, 406, 922, 176, 723, 723, 215, 920, 146, 313, 903, 661, 44, 687, 598, 651, 918, 159, 626, 80, 324, 445, 376, 801, 984, 526, 694, 374, 821, 184, 405, 301, 410, 442, 8, 484, 57, 728, 177, 158, 111, 158, 701, 684, 230, 704, 512, 489, 58, 891, 670, 514, 270, 575, 786, 483, 962, 3, 522, 63, 393, 108, 3, 154, 211, 444, 721, 371, 825, 39, 49, 998, 401, 721, 197, 685, 877, 417, 28, 363, 221, 353, 998, 507, 261, 229, 642, 605, 425, 956, 362, 676, 712, 293, 979, 870, 948, 356, 27, 769, 94, 120, 11, 284, 813, 858, 401, 443, 994, 633, 831, 367, 414, 183, 408, 340, 39, 137, 742, 217, 145, 481, 935, 711, 79, 90, 818, 684, 319, 542, 432, 296, 435, 386, 414, 978, 565, 411, 443, 193, 524, 58, 925, 971, 163, 297, 467, 670, 148, 553, 443, 864, 837, 883, 662, 553, 9, 937, 58, 189, 164, 907, 873, 246, 653, 100, 938, 725, 320, 763, 530, 194, 990, 156, 277, 976, 249, 768, 20, 170, 903, 591, 125, 914, 79, 766, 900, 325, 353, 471, 869, 256, 85, 517, 252, 409, 993, 65, 190, 835, 959, 958, 27, 199, 979, 397, 751, 537, 978, 233, 703, 46, 780, 560, 424, 714, 647, 219, 335, 294, 872, 94, 521, 26, 596, 898, 591, 634, 984, 611, 276, 435, 782, 453, 785, 786, 142, 852, 317, 621, 683, 541, 932, 60, 711, 913, 3, 728, 180, 118, 772, 861, 163, 362, 445, 207, 517, 243, 122, 779, 621, 695, 885, 557, 475, 996, 694, 184, 850, 724, 546, 608, 799, 886, 302, 163, 599, 810, 736, 119, 159, 452, 73, 134, 530, 902, 406, 71, 434, 402, 101, 921, 473, 184, 366, 235, 858, 554, 566, 44, 622, 75, 555, 194, 239, 316, 36, 738, 175, 287, 570, 432, 686, 382, 858, 822, 411, 840, 818, 481, 458, 448, 528, 965, 784, 256, 370, 334, 541, 132, 488, 464, 669, 770, 444, 553, 428, 712, 660, 780, 502, 662, 649, 750, 162, 669, 412, 704, 996, 420, 676, 928, 180, 924, 991, 985, 317, 84, 999, 768, 716, 239, 835, 997, 649, 487, 486, 243, 278, 768, 9, 280, 419, 985, 570, 19, 826, 801, 814, 281, 788, 181, 557, 478, 277, 467, 922, 966, 289, 508, 9, 663, 129, 439, 281, 956, 135, 733, 176, 993, 675, 223, 404, 534, 361, 719, 215, 449, 218, 664, 223, 228, 691, 907, 643, 707, 645, 995, 523, 594, 853, 277, 861, 338, 137, 793, 270, 547, 396, 576, 78, 393, 861, 161, 752, 218, 387, 935, 691, 715, 653, 341, 486, 100, 378, 122, 548, 585, 91, 307, 828, 220, 382, 72, 766, 209, 593, 758, 807, 956, 483, 13, 150, 953, 75, 879, 203, 572, 12, 948, 714, 904, 151, 899, 856, 299, 923, 413, 148, 31, 868, 699, 346, 910, 182, 559, 139, 56, 256, 658, 398, 684, 413, 614, 741, 566, 341, 611, 287, 467, 238, 7, 109, 413, 933, 634, 493, 344, 683, 6, 607, 537, 152, 638, 383, 447, 132, 466, 720, 750, 901, 33, 553, 470, 330, 192, 136, 936, 152, 772, 45, 355, 195, 969, 307, 568, 33, 489, 24, 584, 333, 274, 494, 603, 329, 710, 355, 837, 304, 157, 410, 149, 339, 485, 978, 988, 456, 831, 996, 188, 299, 719, 46, 148, 625, 662, 816, 647, 192, 111, 942, 826, 233, 483, 906, 826, 627, 795, 809, 138, 416, 1000, 233, 634, 177, 473, 425, 580, 85, 937, 213, 441, 29, 735, 357, 985, 766, 459, 119, 391, 514, 586, 957, 948, 851, 544, 416, 2, 248, 959, 13, 297, 357, 875, 994, 60, 604, 319, 827, 823, 550, 101, 849, 788, 696, 310, 59, 490, 616, 225, 818, 628, 532, 343, 7, 925, 775, 487, 388, 840, 677, 609, 506, 978, 883, 934, 443, 578, 183, 954, 155, 459, 210, 872, 171, 124, 566, 465, 963, 22, 864, 967, 222, 524, 182, 821, 898, 807, 722, 570, 381, 195, 373, 963, 552, 468, 36, 11, 735, 335, 931, 487, 514, 510, 529, 450, 82, 290, 872, 957, 868, 481, 453, 7, 949, 80, 209, 517, 659, 559, 98, 275, 310, 483, 236, 306, 385, 708, 367, 950, 466, 447, 15, 374, 953, 122, 189, 838, 47, 841, 885, 964, 512, 222, 609, 325, 504, 995, 242, 194, 119, 769, 619, 299, 200, 926, 159, 308, 388, 455, 688, 161, 781, 894, 467, 449, 302, 886, 958, 339, 479, 284, 20, 712, 257, 671, 854, 532, 404, 378, 951, 738, 677, 468, 921, 915, 956, 166, 126, 158, 862, 82, 812, 126, 35, 771, 837, 195, 957, 242, 417, 35, 558, 734, 94, 242, 57, 955, 687, 463, 186, 172, 659, 824, 660, 384, 867, 883, 948, 682, 667, 620, 717, 601, 158, 467, 122, 312, 955, 754, 30, 449, 95, 241, 625, 513, 421, 58, 592, 963, 376, 421, 313, 171, 224, 679, 168, 1, 469, 677, 332, 653, 866, 558, 501, 243, 466, 207, 489, 417, 147, 528, 85, 599, 420, 716, 71, 660, 459, 41, 641, 485, 597, 353, 729, 501, 914, 657, 659, 236, 495, 474, 285, 121, 773, 793, 512, 764, 607, 308, 200, 835, 699, 344, 834, 265, 940, 409, 163, 463, 659, 854, 207, 540, 173, 170, 871, 664, 318, 376, 380, 369, 333, 212, 137, 775, 26, 467, 597, 27, 801, 134, 769, 758, 177, 13, 243, 16, 825, 910, 659, 344, 955, 894, 1000, 148, 530, 375, 26, 455, 988, 33, 764, 658, 720, 264, 393, 330, 969, 424, 356, 971, 231, 989, 820, 844, 15, 868, 711, 556, 316, 153, 482, 371, 640, 712, 131, 233, 214, 381, 271, 529, 915, 307, 831, 151, 608, 797, 337, 3, 504, 286, 644, 787, 314, 885, 571, 300, 844, 174, 397, 587, 782, 722, 799, 813, 661, 672, 733, 275, 233, 82, 598, 847, 52, 61, 958, 301, 193, 350, 561, 114, 856, 215, 432, 221, 406, 414, 923, 538, 360, 151, 687, 287, 364, 314, 439, 403, 585, 570, 892, 827, 439, 855, 991, 141, 385, 749, 961, 390, 530, 999, 641, 140, 158, 404, 334, 81, 553, 636, 751, 907, 800, 918, 296, 405, 10, 175, 463, 560, 789, 581, 42, 43, 568, 358, 117, 698, 561, 592, 21, 682, 762, 5, 245, 604, 828, 332, 488, 991, 592, 231, 562, 166, 960, 593, 228, 166, 387, 194, 387, 899, 170, 748, 898, 627, 957, 406, 295, 286, 911, 563, 348, 166, 84, 268, 130, 591, 861, 783, 46, 593, 425, 995, 524, 522, 623, 8, 736, 768, 384, 439, 924, 513, 132, 594, 664, 513, 363, 298, 207, 141, 981, 181, 575, 325, 670, 769, 268, 95, 796, 1000, 751, 138, 951, 609, 117, 705, 753, 767, 682, 387, 422, 647, 760, 381, 121, 746, 168, 643, 81, 350, 529, 377, 782, 52, 348, 43, 950, 156, 911, 151, 329, 145, 869, 922, 151, 135, 710, 740, 467, 132, 593, 507, 847, 595, 77, 364, 3, 367, 661, 712, 985, 445, 58, 977, 908, 727, 458, 738, 484, 906, 482, 705, 112, 551, 601, 921, 70, 319, 483, 388, 238, 693, 771, 55, 844, 937, 430, 166, 874, 671, 158, 139, 326, 610, 762, 344, 364, 408, 580, 442, 909, 472, 127, 637, 126, 216, 550, 54, 100, 278, 835, 931, 445, 685, 561, 647, 1, 724, 316, 151, 93, 848, 286, 681, 235, 367, 873, 306, 701, 681, 490, 110, 81, 919, 221, 648, 915, 390, 594, 10, 704, 920, 488, 736, 512, 952, 674, 101, 499, 893, 685, 573, 595, 946, 881, 942, 860, 910, 868, 541, 384, 178, 798, 345, 165, 344, 306, 549, 135, 551, 320, 911, 853, 502, 647, 94, 859, 971, 946, 381, 775, 616, 608, 598, 738, 136, 935, 745, 532, 149, 639, 956, 406, 760, 911, 511, 727, 402, 306, 429, 906, 816, 222, 936, 848, 154, 52, 539, 293, 483, 141, 401, 835, 568, 905, 19, 821, 279, 398, 436, 693, 219, 651, 435, 64, 273, 531, 188, 288, 807, 164, 713, 460, 97, 321, 854, 429, 389, 205, 322, 860, 488, 408, 854, 538, 46, 989, 305, 309, 865, 856, 605, 549, 1000, 505, 358, 878, 985, 760, 303, 584, 509, 271, 797, 655, 385, 305, 169, 878, 923, 896, 978, 928, 961, 309, 988, 908, 926, 770, 999, 285, 395, 2, 808, 784, 106, 99, 109, 11, 506, 900, 628, 757, 355, 77, 339, 805, 595, 8, 378, 64, 814, 844, 628, 162, 969, 718, 564, 64, 497, 848, 75, 23, 58, 282, 877, 817, 519, 93, 937, 708, 924, 273, 286, 965, 705, 938, 454, 328, 257, 940, 496, 826, 853, 1000, 219, 294, 174, 911, 528, 384, 427, 551, 990, 265, 480, 800, 702, 716, 188, 535, 109, 444, 540, 451, 57, 889, 653, 109, 807, 860, 433, 885, 129, 894, 119, 467, 726, 520, 51, 180, 60, 481, 333, 556, 299, 680, 5, 538, 499, 784, 241, 276, 70, 880, 131, 742, 481, 690, 799, 765, 830, 390, 241, 377, 446, 143, 223, 467, 861, 972, 783, 737, 863, 843, 202, 604, 112, 441, 432, 187, 911, 731, 61, 304, 603, 70, 663, 806, 446, 869, 147, 669, 924, 21, 715, 275, 356, 448, 630, 718, 49, 613, 821, 440, 126, 494, 510, 346, 162, 203, 617, 147, 15, 649, 159, 867, 710, 268, 737, 892, 335, 334, 237, 23, 842, 793, 673, 602, 503, 460, 99, 674, 302, 330, 336, 588, 872, 767, 200, 952, 79, 375, 51, 52, 440, 148, 815, 28, 741, 424, 375, 756, 380, 298, 267, 840, 1000, 115, 122, 82, 648, 620, 961, 117, 47, 520, 697, 731, 633, 276, 869, 705, 929, 811, 584, 100, 842, 907, 263, 73, 909, 533, 617, 671, 140, 902, 208, 999, 703, 890, 339, 565, 146, 823, 463, 262, 592, 322, 819, 548, 281, 814, 641, 814, 740, 830, 530, 365, 162, 249, 114, 488, 433, 36, 495, 669, 887, 98, 827, 231, 813, 252, 227, 868, 293, 909, 664, 699, 424, 998, 812, 75, 728, 240, 746, 958, 882, 670, 977, 959, 948, 995, 587, 923, 281, 10, 842, 70, 112, 59, 153, 100, 808, 167, 286, 644, 642, 117, 34, 617, 239, 774, 307, 208, 147, 53, 410, 912, 313, 657, 199, 336, 783, 17, 427, 284, 323, 115, 511, 336, 6, 547, 961, 432, 179, 411, 615, 258, 935, 627, 681, 332, 509, 584, 136, 475, 440, 69, 482, 413, 913, 914, 728, 571, 726, 352, 414, 346, 517, 701, 215, 834, 537, 122, 503, 386, 136, 596, 947, 498, 329, 398, 806, 699, 586, 668, 642, 475, 476, 490, 570, 829, 82, 845, 304, 716, 248, 472, 423, 163, 159, 204, 808, 762, 309, 364, 907, 982, 456, 654, 809, 911, 311, 483, 387, 527, 86, 771, 902, 275, 942, 517, 757, 645, 423, 400, 624, 724, 763, 165, 879, 428, 975, 728, 133, 337, 971, 426, 138, 378, 912, 709, 774, 165, 568, 332, 238, 202, 34, 936, 18, 960, 528, 932, 656, 643, 259, 267, 755, 835, 447, 176, 234, 599, 965, 130, 633, 286, 850, 232, 945, 10, 321, 49, 413, 300, 117, 679, 640, 380, 640, 460, 884, 251, 575, 287, 286, 840, 980, 521, 979, 88, 395, 982, 19, 633, 313, 675, 770, 305, 765, 823, 410, 770, 463, 975, 782, 50, 342, 688, 40, 450, 173, 286, 87, 466, 511, 897, 242, 727, 707, 563, 42, 89, 53, 668, 139, 430, 643, 497, 509, 163, 308, 103, 485, 485, 15, 41, 481, 631, 868, 43, 768, 303, 705, 265, 948, 80, 905, 564, 547, 668, 984, 124, 146, 670, 14, 591, 145, 718, 723, 784, 343, 994, 682, 998, 493, 916, 92, 351, 428, 581, 402, 149, 447, 816, 872, 19, 568, 240, 24, 559, 214, 648, 938, 4, 149, 882, 389, 215, 241, 500, 481, 500, 293, 1000, 560, 654, 26, 689, 937, 510, 47, 98, 110, 616, 361, 879, 120, 775, 800, 726, 273, 315, 558, 996, 49, 78, 279, 752, 485, 868, 422, 714, 965, 623, 585, 385, 594, 363, 345, 337, 324, 472, 928, 200, 591, 297, 407, 949, 342, 293, 841, 76, 743, 758, 285, 770, 547, 961, 804, 451, 345, 746, 165, 495, 190, 762, 242, 883, 761, 789, 268, 416, 287, 988, 58, 50, 229, 598, 848, 62, 288, 661, 132, 868, 45, 855, 941, 351, 705, 46, 474, 742, 433, 219, 950, 311, 420, 508, 916, 392, 450, 869, 501, 570, 718, 530, 425, 779, 576, 173, 98, 920, 814, 689, 718, 229, 147, 160, 708, 937, 981, 522, 436, 124, 249, 872, 231, 589, 803, 874, 153, 316, 114, 251, 377, 104, 186, 363, 64, 230, 859, 273, 180, 131, 224, 852, 433, 457, 764, 229, 166, 951, 514, 239, 931, 58, 697, 641, 370, 46, 576, 50, 231, 733, 707, 386, 866, 665, 329, 233, 596, 440, 501, 337, 570, 224, 242, 443, 519, 714, 572, 366, 261, 140, 574, 358, 977, 574, 511, 112, 384, 654, 74, 131, 488, 645, 587, 405, 587, 388, 571, 993, 924, 181, 492, 431, 973, 564, 63, 647, 30, 454, 964, 668, 170, 951, 433, 877, 395, 709, 739, 872, 481, 144, 560, 703, 438, 537, 331, 361, 943, 770, 768, 257, 710, 138, 391, 964, 567, 344, 107, 322, 184, 285, 901, 340, 209, 487, 749, 558, 407, 379, 940, 747, 47, 512, 724, 292, 438, 261, 412, 733, 596, 421, 841, 64, 483, 332, 266, 874, 680, 603, 511, 975, 724, 907, 84, 996, 261, 651, 347, 277, 612, 855, 93, 735, 456, 895, 153, 690, 1000, 958, 725, 207, 511, 435, 755, 223, 183, 88, 156, 254, 75, 264, 466, 142, 144, 882, 369, 306, 330, 681, 317, 917, 354, 70, 403, 702, 332, 703, 35, 332, 313, 700, 207, 619, 872, 634, 505, 986, 880, 698, 988, 570, 654, 62, 513, 720, 719, 772, 283, 710, 97, 245, 486, 686, 516, 652, 808, 471, 288, 89, 849, 178, 799, 754, 35, 234, 603, 274, 100, 228, 572, 22, 966, 47, 284, 307, 666, 981, 293, 126, 459, 561, 897, 947, 666, 263, 109, 818, 219, 476, 880, 683, 461, 616, 626, 812, 940, 364, 879, 840, 363, 992, 215, 930, 972, 253, 668, 838, 862, 65, 61, 764, 892, 57, 338, 293, 870, 646, 396, 27, 118, 26, 929, 959, 448, 360, 199, 993, 545, 252, 695, 841, 1, 656, 40, 381, 349, 549, 310, 977, 211, 989, 934, 489, 46, 132, 502, 440, 82, 990, 978, 817, 706, 803, 202, 106, 306, 252, 555, 367, 833, 170, 537, 574, 199, 929, 528, 261, 292, 567, 548, 930, 780, 323, 869, 528, 462, 964, 954, 262, 316, 601, 230, 519, 754, 717, 295, 475, 648, 366, 41, 502, 142, 649, 341, 719, 290, 337, 383, 513, 176, 297, 130, 543, 875, 763, 252, 36, 360, 723, 828, 803, 964, 208, 10, 557, 795, 852, 239, 864, 319, 778, 734, 826, 24, 829, 618, 493, 136, 928, 521, 518, 530, 582, 697, 618, 663, 537, 813, 59, 741, 116, 232, 660, 769, 418, 322, 130, 187, 551, 338, 10, 737, 707, 956, 906, 507, 127, 865, 628, 749, 253, 25, 563, 775, 51, 691, 520, 845, 160, 806, 663, 44, 65, 152, 254, 891, 916, 376, 947, 96, 946, 290, 323, 388, 995, 242, 239, 663, 826, 365, 137, 96, 212, 621, 912, 636, 147, 960, 153, 30, 135, 37, 422, 348, 778, 633, 240, 517, 334, 680, 628, 454, 896, 160, 349, 633, 957, 475, 360, 190, 506, 42, 432, 370, 945, 565, 542, 658, 696, 736, 66, 379, 550, 978, 38, 956, 382, 576, 675, 278, 734, 142, 355, 218, 673, 345, 821, 826, 775, 955, 873, 346, 954, 5, 807, 979, 332, 399, 685, 950, 129, 862, 770, 461, 876, 881, 765, 35, 224, 514, 150, 115, 701, 741, 475, 860, 375, 990, 304, 666, 649, 222, 803, 951, 267, 869, 763, 80, 850, 396, 165, 355, 548, 229, 746, 993, 342, 468, 602, 30, 343, 698, 243, 984, 905, 476, 224, 747, 677, 857, 652, 609, 5, 660, 12, 208, 921, 828, 936, 441, 755, 366, 179, 568, 60, 380, 500, 47, 840, 581, 714, 53, 429, 159, 647, 349, 278, 496, 183, 693, 246, 441, 448, 210, 672, 346, 311, 115, 647, 399, 291, 498, 931, 605, 50, 637, 616, 566, 527, 759, 632, 32, 781, 214, 696, 803, 731, 940, 299, 720, 263, 20, 932, 868, 217, 880, 161, 852, 240, 87, 51, 867, 883, 333, 208, 729, 729, 285, 739, 823, 398, 187, 444, 395, 201, 611, 312, 451, 888, 379, 824, 295, 722, 982, 573, 400, 653, 247, 796, 96, 442, 986, 935, 44, 683, 107, 398, 713, 12, 311, 522, 148, 693, 63, 641, 941, 581, 87, 214, 927, 120, 196, 507, 90, 297, 517, 521, 144, 820, 870, 928, 814, 729, 633, 866, 889, 816, 357, 505, 991, 299, 69, 919, 788, 322, 722, 815, 917, 930, 894, 172, 411, 64, 799, 468, 353, 352, 119, 920, 668, 592, 54, 883, 176, 306, 409, 398, 845, 674, 999, 676, 244, 894, 605, 145, 749, 186, 126, 250, 993, 967, 516, 445, 349, 857, 934, 410, 906, 468, 109, 350, 538, 62, 244, 544, 511, 196, 658, 546, 848, 167, 524, 109, 365, 430, 197, 931, 968, 508, 353, 518, 84, 524, 442, 71, 60, 751, 214, 533, 627, 951, 324, 276, 307, 538, 961, 250, 106, 126, 702, 47, 146, 799, 160, 99, 71, 325, 403, 716, 959, 401, 28, 344, 15, 491, 712, 928, 279, 669, 78, 425, 420, 118, 581, 969, 230, 860, 200, 105, 768, 689, 289, 461, 433, 508, 688, 142, 180, 452, 65, 503, 910, 16, 461, 250, 926, 596, 47, 519, 214, 189, 724, 486, 64, 618, 672, 911, 870, 669, 519, 775, 713, 570, 242, 870, 935, 740, 74, 200, 123, 252, 966, 224, 382, 596, 464, 274, 343, 560, 464, 763, 322, 238, 806, 652, 196, 896, 69, 511, 973, 964, 643, 130, 126, 925, 919, 642, 750, 655, 891, 605, 128, 516, 562, 901, 965, 482, 665, 991, 235, 416, 876, 336, 904, 236, 725, 919, 70, 614, 240, 414, 117, 210, 252, 265, 396, 983, 700, 211, 873, 335, 49, 838, 963, 608, 120, 195, 327, 734, 456, 205, 229, 723, 687, 301, 440, 491, 115, 198, 217, 566, 275, 223, 701, 724, 27, 235, 445, 264, 757, 369, 3, 766, 921, 553, 401, 516, 16, 17, 891, 346, 187, 474, 622, 493, 117, 314, 431, 283, 881, 609, 848, 300, 570, 183, 359, 409, 243, 411, 290, 138, 439, 897, 927, 821, 418, 529, 683, 240, 697, 441, 569, 870, 209, 333, 940, 688, 615, 887, 82, 149, 997, 432, 817, 66, 127, 442, 187, 598, 567, 943, 853, 921, 483, 892, 393, 244, 894, 276, 466, 175, 306, 363, 51, 641, 844, 368, 565, 471, 980, 800, 395, 64, 273, 241, 10, 1, 335, 584, 832, 354, 307, 412, 137, 841, 657, 963, 864, 231, 280, 483, 373, 366, 830, 48, 596, 84, 796, 126, 3, 261, 152, 383, 963, 413, 276, 714, 443, 543, 107, 893, 166, 167, 6, 180, 431, 876, 620, 319, 380, 870, 934, 133, 565, 687, 567, 471, 185, 726, 626, 336, 440, 475, 992, 719, 37, 739, 905, 591, 712, 576, 81, 190, 365, 344, 179, 963, 281, 386, 95, 55, 281, 219, 359, 299, 52, 11, 612, 471, 542, 275, 653, 158, 951, 700, 575, 56, 847, 700, 650, 486, 560, 564, 114, 609, 145, 396, 524, 793, 188, 11, 845, 537, 964, 127, 889, 434, 633, 5, 699, 199, 338, 679, 5, 171, 453, 905, 679, 998, 93, 750, 404, 509, 132, 679, 869, 302, 930, 290, 835, 110, 726, 324, 305, 377, 61, 95, 409, 563, 34, 230, 4, 862, 460, 658, 43, 434, 109, 174, 560, 253, 480, 274, 263, 147, 3, 177, 417, 907, 591, 378, 586, 645, 812, 663, 192, 688, 173, 649, 104, 690, 20, 518, 831, 352, 23, 100, 61, 300, 170, 147, 221, 821, 576, 824, 132, 677, 815, 531, 817, 735, 745, 10, 285, 900, 214, 595, 424, 84, 385, 333, 524, 930, 360, 461, 288, 963, 187, 464, 616, 79, 812, 215, 969, 955, 408, 780, 512, 646, 417, 202, 13, 631, 111, 204, 618, 149, 257, 469, 874, 844, 73, 795, 839, 250, 942, 404, 831, 376, 58, 615, 502, 320, 97, 861, 8, 403, 330, 348, 989, 665, 156, 696, 362, 947, 685, 695, 603, 974, 482, 947, 358, 531, 815, 643, 386, 636, 661, 179, 98, 326, 659, 626, 734, 377, 306, 278, 305, 270, 243, 84, 371, 247, 920, 71, 824, 933, 283, 12, 138, 962, 157, 901, 424, 362, 441, 975, 378, 774, 119, 13, 7, 953, 855, 58, 507, 253, 680, 532, 690, 335, 259, 458, 265, 1000, 716, 26, 986, 590, 675, 39, 147, 396, 651, 592, 15, 633, 629, 65, 635, 50, 414, 465, 967, 227, 418, 768, 826, 535, 428, 586, 134, 939, 266, 736, 90, 584, 218, 527, 256, 425, 148, 336, 441, 831, 433, 384, 364, 353, 947, 301, 817, 157, 478, 879, 12, 583, 820, 677, 737, 672, 944, 619, 327, 85, 464, 519, 672, 431, 868, 84, 36, 870, 172, 322, 925, 266, 944, 900, 958, 311, 423, 401, 476, 305, 30, 22, 17, 446, 992, 245, 451, 886, 741, 627, 925, 963, 63, 726, 823, 634, 723, 423, 634, 983, 278, 640, 990, 840, 361, 688, 790, 883, 594, 530, 774, 220, 253, 982, 985, 314, 855, 712, 682, 240, 224, 130, 938, 195, 872, 262, 207, 464, 401, 705, 259, 248, 293, 855, 807, 606, 493, 213, 127, 959, 259, 839, 393, 912, 391, 477, 766, 7, 462, 225, 543, 21, 517, 266, 549, 524, 936, 403, 895, 340, 607, 698, 124, 162, 520, 528, 988, 284, 749, 622, 352, 98, 703, 570, 952, 606, 574, 484, 13, 108, 723, 209, 577, 248, 225, 108, 30, 90, 84, 8, 250, 652, 190, 122, 152, 636, 623, 796, 347, 776, 754, 47, 112, 467, 852, 385, 821, 387, 655, 663, 291, 910, 432, 606, 253, 687, 325, 287, 466, 939, 422, 156, 277, 585, 934, 944, 620, 440, 712, 62, 960, 7, 399, 284, 111, 809, 723, 128, 453, 51, 865, 669, 484, 146, 44, 454, 884, 728, 619, 337, 618, 703, 912, 357, 726, 60, 775, 796, 141, 843, 596, 540, 694, 271, 618, 255, 182, 976, 424, 216, 786, 578, 949, 546, 712, 430, 145, 584, 819, 233, 455, 652, 184, 27, 562, 366, 980, 663, 132, 424, 855, 818, 726, 694, 843, 760, 445, 505, 300, 881, 969, 968, 724, 803, 365, 477, 770, 27, 622, 86, 349, 331, 990, 32, 213, 379, 105, 408, 944, 706, 407, 881, 104, 366, 677, 744, 510, 116, 36, 903, 568, 847, 125, 985, 257, 912, 843, 399, 38, 270, 332, 85, 665, 822, 239, 366, 750, 161, 987, 438, 694, 951, 555, 158, 510, 624, 132, 901, 268, 792, 306, 106, 418, 778, 935, 11, 227, 946, 259, 151, 682, 34, 343, 817, 111, 143, 593, 59, 220, 35, 17, 763, 269, 278, 652, 38, 775, 395, 885, 608, 906, 164, 35, 336, 203, 139, 987, 876, 548, 155, 570, 306, 593, 709, 415, 381, 847, 881, 894, 138, 560, 647, 685, 45, 658, 616, 710, 62, 245, 153, 814, 589, 205, 930, 525, 414, 253, 517, 821, 912, 513, 510, 207, 101, 446, 588, 74, 493, 417, 566, 838, 364, 848, 234, 365, 315, 724, 153, 239, 67, 219, 969, 727, 927, 661, 3, 407, 309, 72, 847, 321, 212, 45, 807, 829, 475, 53, 793, 983, 490, 901, 787, 38, 596, 910, 220, 867, 266, 862, 257, 661, 803, 252, 772, 260, 46, 584, 877, 431, 337, 665, 921, 691, 297, 282, 908, 622, 783, 433, 743, 10, 505, 912, 509, 775, 421, 418, 810, 144, 724, 700, 59, 241, 546, 246, 628, 519, 563, 904, 607, 188, 328, 522, 958, 90, 555, 903, 795, 434, 425, 376, 699, 69, 556, 788, 145, 185, 684, 286, 400, 599, 259, 436, 493, 60, 320, 145, 152, 359, 918, 3, 635, 576, 574, 901, 971, 463, 388, 973, 891, 833, 431, 548, 348, 328, 6, 362, 367, 73, 843, 128, 928, 414, 307, 872, 841, 685, 518, 914, 572, 195, 839, 202, 546, 826, 104, 708, 711, 28, 783, 747, 697, 241, 177, 246, 154, 679, 373, 318, 845, 410, 626, 434, 217, 848, 532, 424, 200, 366, 280, 307, 847, 7, 78, 340, 554, 864, 951, 708, 387, 786, 821, 102, 121, 425, 448, 528, 652, 662, 316, 272, 909, 22, 429, 634, 456, 278, 867, 498, 26, 397, 665, 99, 864, 934, 455, 754, 269, 6, 639, 316, 814, 820, 59, 786, 48, 9, 577, 960, 917, 968, 569, 6, 256, 755, 388, 22, 813, 813, 615, 536, 273, 588, 911, 395, 284, 111, 911, 690, 364, 597, 855, 951, 835, 328, 90, 12, 57, 898, 280, 893, 863, 593, 922, 613, 440, 773, 536, 556, 560, 71, 728, 480, 253, 509, 588, 106, 772, 291, 970, 714, 159, 669, 390, 452, 23, 692, 694, 106, 101, 291, 494, 501, 361, 453, 953, 446, 354, 960, 363, 978, 616, 850, 197, 226, 992, 580, 701, 828, 878, 846, 373, 296, 321, 469, 133, 902, 172, 708, 877, 908, 718, 625, 509, 795, 330, 814, 254, 718, 993, 842, 772, 153, 171, 919, 883, 89, 777, 966, 950, 312, 643, 538, 644, 352, 392, 533, 474, 864, 554, 758, 422, 114, 309, 119, 330, 487, 32, 112, 759, 644, 891, 653, 808, 66, 797, 292, 652, 364, 713, 759, 620, 699, 361, 78, 677, 346, 332, 403, 112, 737, 59, 854, 327, 987, 865, 157, 911, 817, 461, 701, 796, 707, 425, 105, 221, 523, 647, 266, 710, 423, 884, 409, 255, 560, 503, 668, 873, 343, 352, 913, 178, 677, 639, 950, 567, 386, 759, 784, 436, 781, 360, 980, 416, 822, 598, 824, 35, 706, 274, 13, 82, 259, 779, 888, 816, 291, 413, 627, 792, 65, 927, 363, 241, 758, 371, 762, 195, 927, 431, 211, 341, 380, 480, 883, 821, 53, 226, 119, 743, 852, 289, 692, 738, 236, 87, 171, 589, 949, 262, 18, 493, 923, 440, 374, 369, 559, 738, 317, 931, 128, 84, 319, 651, 255, 376, 67, 267, 252, 171, 915, 168, 714, 581, 39, 480, 962, 914, 577, 757, 964, 789, 360, 604, 689, 399, 785, 438, 954, 619, 624, 314, 630, 324, 215, 632, 366, 450, 742, 244, 949, 786, 64, 747, 632, 93, 381, 744, 754, 200, 661, 518, 538, 564, 869, 64, 231, 44, 400, 305, 615, 808, 131, 52, 260, 47, 182, 867, 687, 971, 334, 288, 379, 935, 767, 422, 46, 351, 139, 141, 644, 677, 771, 456, 108, 309, 347, 526, 539, 341, 69, 548, 671, 337, 777, 511, 703, 758, 198, 225, 240, 555, 235, 353, 138, 1000, 962, 641, 683, 299, 646, 871, 460, 294, 939, 85, 420, 522, 128, 222, 220, 706, 988, 999, 181, 463, 789, 935, 836, 71, 244, 163, 753, 718, 312, 28, 677, 632, 172, 856, 101, 292, 58, 717, 195, 523, 515, 690, 859, 352, 533, 29, 902, 480, 200, 146, 947, 475, 471, 5, 196, 504, 743, 409, 208, 696, 879, 497, 479, 269, 482, 949, 199, 15, 467, 805, 160, 75, 35, 508, 484, 983, 689, 470, 53, 143, 660, 483, 164, 167, 656, 537, 197, 519, 882, 719, 70, 276, 915, 898, 92, 573, 618, 233, 914, 501, 221, 828, 161, 524, 648, 463, 452, 200, 829, 237, 756, 498, 280, 313, 772, 328, 788, 50, 980, 601, 253, 15, 535, 706, 328, 207, 208, 492, 280, 938, 229, 415, 911, 550, 380, 450, 922, 44, 570, 447, 441, 191, 231, 362, 748, 960, 787, 260, 915, 997, 633, 57, 869, 538, 109, 635, 151, 750, 782, 671, 938, 728, 150, 867, 977, 91, 829, 51, 119, 909, 48, 682, 772, 900, 8, 729, 958, 75, 752, 543, 766, 422, 659, 602, 732, 447, 892, 731, 866, 605, 796, 370, 193, 338, 162, 774, 308, 131, 109, 467, 689, 101, 709, 527, 532, 777, 393, 422, 964, 815, 315, 775, 857, 552, 70, 664, 348, 766, 843, 205, 739, 797, 829, 907, 954, 86, 547, 243, 917, 696, 993, 34, 519, 485, 629, 472, 837, 129, 995, 532, 193, 128, 546, 537, 332, 392, 585, 464, 299, 816, 761, 678, 855, 578, 889, 719, 109, 733, 871, 324, 416, 479, 735, 888, 96, 201, 892, 639, 57, 810, 377, 477, 49, 108, 294, 20, 913, 846, 187, 239, 271, 413, 340, 778, 621, 710, 970, 627, 87, 861, 62, 32, 894, 444, 483, 290, 219, 843, 434, 885, 103, 277, 265, 23, 57, 461, 412, 333, 364, 152, 171, 744, 969, 263, 340, 909, 223, 112, 69, 824, 626, 788, 415, 589, 154, 672, 588, 638, 215, 685, 948, 576, 718, 765, 25, 154, 933, 799, 789, 379, 820, 39, 523, 758, 986, 838, 532, 917, 33, 915, 769, 314, 497, 905, 925, 714, 226, 950, 990, 250, 566, 69, 950, 694, 25, 786, 196, 891, 614, 341, 675, 28, 571, 602, 641, 189, 177, 844, 511, 787, 933, 726, 338, 411, 425, 838, 187, 247, 488, 118, 380, 876, 83, 944, 739, 313, 327, 770, 351, 275, 43, 117, 678, 488, 77, 987, 73, 206, 195, 209, 851, 408, 450, 189, 343, 269, 623, 766, 68, 194, 576, 370, 462, 68, 219, 335, 696, 406, 391, 430, 662, 384, 608, 938, 778, 228, 997, 123, 58, 639, 542, 534, 190, 121, 832, 544, 921, 572, 226, 174, 320, 969, 897, 610, 571, 625, 766, 877, 560, 17, 188, 210, 729, 291, 805, 81, 172, 462, 504, 1000, 449, 770, 193, 75, 609, 873, 129, 551, 724, 977, 59, 259, 790, 54, 534, 231, 745, 415, 24, 854, 50, 803, 895, 409, 755, 453, 201, 503, 56, 819, 703, 857, 223, 50, 487, 935, 508, 241, 99, 942, 891, 19, 536, 54, 323, 943, 512, 871, 565, 124, 810, 930, 409, 64, 540, 162, 166, 908, 344, 514, 450, 738, 812, 768, 397, 786, 279, 772, 487, 208, 294, 873, 44, 770, 491, 389, 412, 178, 374, 470, 696, 106, 227, 175, 474, 821, 820, 707, 625, 992, 783, 952, 937, 985, 204, 191, 20, 5, 71, 774, 977, 902, 851, 911, 138, 292, 418, 141, 442, 140, 108, 783, 228, 124, 582, 416, 368, 116, 2, 349, 381, 351, 429, 805, 381, 129, 90, 425, 713, 620, 305, 299, 708, 109, 670, 113, 575, 65, 984, 794, 138, 210, 982, 754, 739, 844, 947, 367, 41, 661, 699, 762, 93, 415, 137, 74, 149, 867, 97, 36, 820, 985, 57, 930, 897, 941, 176, 499, 587, 111, 178, 795, 86, 609, 410, 726, 148, 694, 460, 167, 652, 126, 510, 449, 676, 753, 697, 927, 862, 583, 224, 105, 834, 338, 548, 175, 71, 174, 125, 365, 674, 812, 514, 504, 790, 416, 205, 324, 165, 317, 956, 343, 596, 576, 734, 88, 928, 128, 886, 772, 792, 529, 697, 109, 60, 555, 214, 595, 608, 809, 197, 270, 762, 206, 612, 880, 96, 427, 549, 163, 762, 433, 856, 999, 917, 680, 693, 906, 840, 416, 917, 378, 492, 878, 223, 33, 219, 143, 626, 256, 463, 377, 840, 17, 460, 547, 565, 230, 3, 696, 69, 625, 687, 73, 524, 550, 389, 1, 440, 775, 524, 61, 920, 991, 910, 298, 493, 565, 364, 368, 186, 793, 925, 842, 915, 417, 259, 93, 568, 183, 915, 154, 81, 564, 779, 424, 548, 579, 894, 869, 576, 505, 519, 251, 375, 98, 724, 616, 591, 753, 563, 899, 643, 200, 910, 803, 883, 543, 67, 479, 932, 542, 741, 627, 18, 793, 908, 306, 134, 749, 436, 499, 473, 494, 687, 259, 200, 326, 145, 504, 104, 375, 520, 342, 254, 923, 415, 508, 481, 519, 221, 128, 375, 22, 507, 932, 669, 364, 511, 962, 357, 64, 691, 409, 920, 534, 331, 672, 920, 447, 485, 479, 363, 117, 709, 619, 735, 327, 46, 902, 777, 50, 939, 856, 736, 13, 753, 281, 799, 9, 643, 577, 208, 491, 225, 110, 577, 242, 265, 907, 229, 145, 225, 757, 634, 885, 407, 958, 320, 763, 216, 890, 342, 634, 842, 121, 50, 56, 339, 215, 959, 607, 491, 25, 679, 439, 13, 109, 503, 120, 473, 290, 78, 561, 652, 688, 342, 10, 675, 545, 687, 875, 617, 560, 844, 705, 779, 189, 64, 610, 81, 759, 430, 576, 966, 925, 972, 961, 500, 885, 304, 472, 688, 886, 482, 610, 923, 993, 32, 543, 100, 90, 304, 866, 717, 296, 875, 21, 693, 335, 405, 914, 788, 845, 343, 594, 249, 792, 516, 418, 340, 393, 436, 361, 356, 657, 643, 658, 372, 971, 143, 697, 123, 122, 682, 974, 961, 513, 607, 872, 186, 797, 719, 610, 97, 791, 264, 283, 409, 458, 618, 24, 127, 657, 548, 901, 214, 952, 354, 539, 770, 853, 570, 470, 433, 747, 814, 609, 386, 111, 693, 819, 909, 682, 896, 426, 294, 161, 902, 353, 540, 839, 355, 835, 446, 576, 727, 753, 594, 240, 947, 924, 893, 693, 949, 793, 962, 178, 424, 998, 495, 367, 581, 560, 814, 472, 344, 308, 158, 397, 910, 782, 566, 519, 883, 714, 38, 487, 446, 145, 305, 689, 55, 60, 159, 596, 216, 602, 451, 159, 204, 866, 574, 49, 980, 182, 309, 638, 707, 472, 289, 783, 757, 275, 949, 780, 681, 554, 381, 299, 445, 793, 758, 679, 745, 146, 704, 271, 967, 329, 746, 672, 382, 411, 134, 262, 532, 694, 593, 700, 620, 323, 782, 572, 846, 129, 420, 640, 680, 523, 269, 989, 316, 187, 461, 85, 330, 125, 858, 87, 209, 414, 737, 409, 586, 872, 359, 724, 783, 700, 107, 120, 365, 757, 657, 779, 639, 600, 923, 738, 361, 169, 798, 894, 921, 681, 902, 632, 350, 849, 745, 187, 808, 171, 540, 261, 947, 86, 133, 46, 466, 808, 134, 930, 848, 337, 558, 313, 448, 437, 648, 409, 631, 936, 51, 589, 908, 91, 633, 6, 860, 640, 949, 815, 475, 189, 34, 879, 706, 231, 816, 124, 281, 107, 730, 785, 321, 48, 537, 866, 103, 922, 187, 817, 952, 807, 423, 852, 934, 887, 812, 184, 944, 223, 548, 710, 641, 582, 374, 960, 221, 877, 93, 374, 83, 934, 490, 854, 734, 847, 262, 852, 948, 562, 592, 652, 224, 418, 623, 842, 463, 311, 590, 371, 843, 258, 341, 529, 113, 443, 427, 889, 704, 433, 855, 217, 187, 665, 347, 337, 803, 915, 345, 623, 269, 847, 14, 342, 27, 202, 116, 970, 96, 202, 590, 552, 120, 288, 625, 83, 541, 340, 465, 418, 312, 52, 235, 456, 47, 292, 561, 2, 924, 724, 98, 159, 287, 142, 773, 327, 208, 206, 361, 907, 576, 982, 772, 109, 989, 703, 194, 477, 755, 645, 812, 828, 737, 812, 751, 336, 270, 46, 602, 749, 230, 759, 14, 998, 827, 916, 335, 903, 764, 170, 174, 387, 23, 756, 484, 703, 648, 358, 685, 116, 244, 79, 730, 627, 650, 357, 263, 991, 619, 742, 715, 708, 491, 232, 471, 573, 415, 268, 113, 582, 378, 690, 340, 170, 21, 918, 606, 493, 846, 248, 705, 516, 399, 109, 940, 720, 677, 880, 447, 712, 52, 428, 37, 243, 831, 711, 771, 476, 140, 916, 262, 701, 745, 845, 954, 531, 597, 810, 37, 481, 662, 437, 972, 201, 625, 77, 1000, 945, 60, 198, 770, 867, 328, 718, 162, 537, 203, 974, 569, 436, 626, 566, 40, 470, 137, 514, 683, 644, 223, 204, 535, 710, 942, 765, 908, 522, 734, 341, 210, 574, 828, 129, 375, 945, 135, 540, 547, 251, 26, 410, 18, 606, 265, 489, 692, 691, 815, 500, 567, 106, 933, 619, 554, 95, 398, 359, 821, 28, 616, 70, 533, 280, 211, 87, 495, 210, 498, 668, 819, 95, 309, 976, 228, 900, 616, 984, 849, 204, 582, 416, 675, 699, 356, 589, 260, 280, 483, 679, 560, 728, 110, 289, 416, 508, 356, 667, 133, 408, 539, 843, 205, 837, 743, 73, 463, 673, 317, 35, 579, 839, 20, 827, 855, 819, 880, 115, 416, 482, 168, 643, 735, 331, 332, 340, 999, 75, 372, 725, 665, 380, 950, 372, 710, 766, 662, 682, 679, 305, 78, 466, 446, 664, 258, 155, 411, 748, 329, 524, 716, 231, 271, 7, 653, 738, 277, 842, 434, 185, 589, 948, 675, 289, 558, 419, 415, 69, 227, 508, 548, 317, 184, 162, 738, 766, 281, 86, 514, 646, 891, 631, 315, 559, 439, 626, 930, 888, 52, 252, 477, 673, 475, 682, 406, 649, 326, 23, 590, 554, 956, 884, 11, 101, 680, 300, 627, 601, 126, 673, 242, 865, 536, 683, 60, 49, 107, 763, 947, 297, 151, 923, 535, 416, 83, 136, 470, 684, 289, 55, 296, 184, 442, 696, 555, 409, 686, 300, 578, 241, 849, 438, 278, 602, 282, 979, 206, 285, 242, 178, 674, 973, 287, 186, 878, 968, 560, 508, 72, 384, 608, 242, 100, 140, 824, 517, 67, 240, 113, 254, 597, 486, 725, 693, 602, 226, 375, 493, 552, 287, 841, 439, 232, 831, 814, 331, 459, 54, 174, 395, 941, 485, 680, 969, 451, 903, 34, 797, 532, 867, 655, 857, 289, 175, 851, 988, 311, 298, 568, 933, 332, 717, 255, 525, 721, 337, 486, 499, 230, 543, 798, 973, 318, 880, 915, 587, 494, 895, 969, 438, 667, 972, 90, 287, 114, 184, 286, 578, 62, 140, 434, 154, 456, 918, 276, 181, 600, 307, 115, 851, 958, 396, 861, 665, 142, 1000, 913, 784, 82, 501, 951, 70, 714, 644, 876, 87, 763, 42, 563, 864, 4, 76, 404, 73, 622, 229, 897, 845, 217, 574, 229, 897, 755, 147, 211, 637, 884, 440, 946, 65, 982, 991, 376, 703, 166, 1000, 688, 564, 771, 628, 293, 181, 26, 126, 26, 216, 501, 149, 270, 670, 234, 372, 19, 323, 4, 4, 996, 600, 247, 719, 98, 853, 733, 280, 419, 624, 381, 956, 408, 340, 785, 672, 111, 241, 117, 754, 235, 912, 550, 630, 908, 782, 567, 261, 919, 239, 340, 899, 60, 816, 524, 198, 62, 467, 350, 521, 240, 540, 310, 499, 962, 20, 663, 837, 642, 48, 558, 756, 942, 756, 216, 263, 399, 558, 993, 641, 815, 364, 967, 763, 778, 615, 890, 417, 681, 704, 824, 632, 73, 806, 876, 249, 955, 650, 420, 118, 348, 439, 991, 623, 371, 97, 565, 430, 470, 373, 576, 756, 538, 526, 332, 706, 567, 230, 697, 470, 494, 8, 33, 922, 772, 168, 316, 739, 453, 375, 508, 30, 415, 328, 977, 470, 610, 880, 8, 210, 24, 765, 557, 136, 810, 415, 802, 253, 471, 142, 959, 599, 154, 957, 27, 629, 20, 391, 324, 365, 223, 211, 418, 172, 388, 73, 498, 783, 897, 1000, 29, 56, 921, 283, 659, 681, 493, 111, 338, 760, 8, 493, 205, 635, 926, 816, 504, 367, 779, 357, 285, 177, 319, 805, 420, 242, 984, 805, 95, 763, 839, 574, 71, 865, 583, 571, 452, 896, 361, 677, 720, 467, 149, 438, 358, 172, 126, 303, 167, 296, 6, 830, 584, 346, 104, 893, 201, 158, 606, 940, 802, 951, 969, 229, 140, 842, 545, 20, 113, 919, 594, 740, 950, 145, 702, 25, 917, 566, 53, 730, 611, 376, 355, 392, 204, 723, 173, 803, 489, 765, 986, 271, 888, 848, 606, 539, 587, 315, 6, 735, 261, 624, 763, 131, 191, 16, 234, 359, 518, 152, 548, 72, 676, 977, 763, 502, 608, 473, 657, 582, 704, 940, 666, 22, 607, 639, 705, 215, 199, 168, 676, 190, 408, 640, 465, 115, 69, 463, 896, 146, 472, 302, 570, 120, 77, 715, 974, 5, 934, 321, 671, 285, 987, 468, 477, 30, 500, 499, 296, 878, 730, 14, 551, 562, 70, 405, 671, 782, 321, 5, 866, 113, 653, 309, 279, 440, 481, 538, 242, 80, 293, 923, 128, 704, 621, 550, 77, 687, 813, 219, 68, 621, 8, 242, 996, 471, 550, 929, 66, 665, 671, 182, 197, 430, 242, 277, 786, 964, 861, 871, 942, 783, 266, 312, 722, 198, 624, 870, 639, 751, 988, 482, 388, 354, 522, 673, 263, 159, 13, 296, 890, 304, 639, 3, 646, 318, 318, 851, 573, 467, 251, 120, 548, 270, 132, 339, 887, 421, 697, 344, 946, 769, 307, 616, 832, 467, 373, 146, 781, 978, 813, 853, 532, 83, 751, 756, 36, 549, 945, 358, 107, 417, 273, 381, 772, 554, 613, 883, 377, 318, 103, 429, 34, 382, 258, 838, 340, 474, 33, 323, 617, 577, 27, 730, 723, 101, 575, 941, 157, 517, 568, 848, 902, 501, 458, 619, 64, 558, 300, 385, 780, 585, 249, 232, 115, 317, 693, 509, 674, 21, 454, 445, 728, 786, 708, 141, 132, 904, 766, 135, 701, 306, 748, 225, 894, 364, 484, 559, 822, 327, 707, 172, 281, 357, 73, 982, 788, 762, 280, 777, 63, 998, 794, 550, 197, 78, 3, 65, 924, 579, 144, 124, 52, 840, 652, 853, 666, 451, 44, 164, 492, 139, 826, 285, 604, 504, 396, 558, 18, 305, 183, 635, 181, 533, 120, 276, 416, 859, 994, 257, 987, 812, 864, 979, 595, 192, 69, 667, 307, 634, 861, 220, 332, 861, 9, 805, 666, 206, 478, 897, 971, 508, 644, 895, 834, 39, 973, 59, 295, 986, 252, 860, 164, 196, 890, 566, 928, 413, 25, 972, 530, 661, 830, 257, 153, 24, 550, 121, 734, 473, 23, 687, 922, 894, 692, 306, 30, 186, 334, 322, 958, 10, 909, 403, 959, 502, 559, 168, 430, 904, 481, 451, 590, 945, 957, 40, 5, 49, 46, 815, 583, 460, 480, 325, 987, 885, 631, 855, 543, 172, 798, 69, 574, 586, 929, 713, 451, 443, 870, 347, 425, 85, 858, 122, 545, 593, 822, 772, 905, 690, 129, 677, 359, 530, 670, 918, 759, 172, 352, 449, 185, 156, 905, 559, 404, 638, 361, 21, 161, 670, 235, 594, 476, 867, 504, 412, 490, 276, 88, 892, 155, 858, 862, 304, 710, 589, 773, 640, 11, 234, 776, 831, 782, 626, 523, 339, 694, 669, 195, 772, 487, 433, 462, 98, 397, 784, 685, 782, 751, 265, 439, 889, 437, 426, 420, 193, 426, 877, 750, 154, 232, 96, 153, 239, 419, 407, 574, 462, 25, 689, 31, 165, 312, 422, 685, 33, 893, 183, 724, 74, 470, 710, 541, 354, 481, 120, 958, 860, 837, 486, 837, 593, 741, 478, 721, 817, 776, 255, 933, 309, 548, 150, 536, 832, 713, 653, 605, 114, 813, 519, 874, 523, 88, 524, 214, 566, 38, 137, 791, 806, 156, 925, 896, 6, 503, 23, 234, 417, 828, 752, 943, 74, 831, 952, 497, 274, 1, 564, 497, 442, 329, 699, 10, 109, 573, 239, 832, 126, 999, 195, 859, 532, 335, 604, 157, 95, 881, 60, 233, 122, 442, 577, 49, 95, 695, 429, 355, 205, 114, 968, 742, 952, 553, 40, 908, 688, 373, 191, 730, 795, 134, 771, 702, 504, 97, 705, 597, 901, 779, 289, 675, 648, 643, 994, 547, 662, 739, 877, 314, 449, 901, 182, 265, 292, 958, 698, 267, 661, 662, 287, 891, 333, 384, 220, 776, 267, 927, 677, 661, 976, 771, 56, 436, 608, 544, 675, 818, 552, 457, 50, 207, 386, 578, 839, 401, 590, 612, 118, 606, 589, 70, 618, 934, 377, 575, 653, 209, 969, 363, 431, 275, 174, 879, 417, 796, 75, 499, 862, 96, 643, 244, 134, 599, 249, 374, 151, 220, 315, 214, 858, 758, 931, 897, 809, 101, 592, 80, 936, 882, 740, 869, 212, 41, 386, 927, 567, 17, 266, 353, 823, 446, 727, 115, 240, 310, 686, 429, 955, 130, 703, 644, 470, 628, 842, 373, 36, 845, 275, 594, 312, 354, 313, 509, 749, 547, 354, 806, 917, 315, 640, 37, 795, 32, 669, 453, 701, 639, 581, 376, 412, 443, 537, 333, 833, 963, 539, 230, 632, 716, 318, 169, 13, 435, 367, 771, 830, 241, 67, 171, 523, 804, 632, 413, 955, 712, 929, 954, 612, 517, 547, 403, 259, 962, 345, 907, 133, 525, 179, 429, 979, 296, 584, 375, 995, 786, 930, 856, 965, 766, 49, 618, 568, 431, 814, 226, 933, 309, 982, 714, 698, 232, 911, 915, 545, 126, 352, 206, 780, 656, 198, 419, 591, 627, 952, 40, 991, 765, 163, 800, 323, 281, 663, 117, 565, 186, 593, 707, 406, 9, 770, 466, 505, 237, 752, 659, 404, 121, 828, 176, 508, 934, 928, 403, 946, 123, 751, 408, 979, 909, 221, 739, 12, 612, 680, 412, 542, 911, 199, 380, 187, 66, 632, 314, 423, 954, 948, 271, 558, 77, 730, 739, 587, 228, 393, 345, 729, 780, 641, 338, 263, 836, 271, 999, 731, 333, 19, 185, 713, 911, 747, 278, 852, 354, 870, 758, 363, 727, 265, 725, 47, 661, 76, 556, 480, 512, 254, 378, 745, 276, 787, 830, 977, 484, 11, 850, 273, 506, 458, 657, 551, 482, 328, 33, 456, 790, 671, 686, 49, 99, 953, 800, 402, 708, 771, 979, 545, 213, 883, 649, 744, 522, 632, 572, 433, 103, 561, 89, 446, 486, 163, 624, 285, 564, 908, 232, 719, 182, 476, 327, 412, 54, 607, 104, 879, 985, 263, 755, 257, 496, 522, 289, 876, 369, 776, 17, 88, 319, 579, 996, 785, 183, 612, 539, 984, 855, 214, 812, 335, 786, 510, 631, 361, 576, 467, 86, 395, 548, 634, 592, 706, 449, 845, 585, 121, 385, 64, 726, 22, 64, 745, 671, 851, 320, 94, 877, 771, 397, 565, 987, 553, 448, 733, 239, 48, 788, 675, 253, 294, 305, 11, 92, 418, 236, 613, 48, 256, 313, 823, 305, 336, 501, 874, 901, 90, 639, 472, 797, 182, 507, 976, 8, 75, 10, 279, 451, 471, 698, 272, 248, 754, 627, 459, 468, 609, 171, 579, 689, 364, 244, 395, 428, 43, 698, 650, 32, 26, 545, 335, 326, 588, 411, 180, 872, 288, 304, 758, 725, 364, 23, 729, 941, 148, 160, 764, 476, 52, 501, 396, 219, 371, 111, 340, 380, 339, 115, 767, 233, 143, 250, 855, 871, 526, 84, 370, 970, 876, 492, 328, 851, 164}, 2719
	fmt.Println(maxScore(cardPoints, k) == 1372863)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func maxScore(cardPoints []int, k int) int {
	ans := 0
	for i := 0; i <= k; i++ {
		l := cardPoints[:i]
		r := cardPoints[len(cardPoints)-k+i:]
		// fmt.Println(l, r)
		sum := 0
		// l = append(l, r...)
		for _, v := range l {
			sum += v
		}
		for _, v := range r {
			sum += v
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}
