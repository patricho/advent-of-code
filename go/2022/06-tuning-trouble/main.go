package tuningtrouble06

import (
	"fmt"
)

/*
test-inputs and expected result
	part 1
		mjqjpqmgbljsphdztnvjfqwrcgsmlb / 7
		bvwbjplbgvbhsrlpgdmjqwftvncz / 5
		nppdvjthqldpwncqszvftbrmjlhg / 6
		nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg / 10
		zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw / 11
	part 2
		mjqjpqmgbljsphdztnvjfqwrcgsmlb / 19
		bvwbjplbgvbhsrlpgdmjqwftvncz / 23
		nppdvjthqldpwncqszvftbrmjlhg / 23
		nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg / 29
		zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw / 26

*/

func Part1() {
	checkMessages(4)
}

func Part2() {
	checkMessages(14)
}

func checkMessages(length int) {
	messages := []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		"cdhccdbdggfjjgssjzjzggjnjpnpbbzbnzzflfjfnfrrpvrvbrvvrvggvlvnnbrnrcncsnndbndbnndbdndfdrdvrvvndvvbggnrrnbrnntffgttwzwnnmvmcvvhsstzzlnlwlttbzzpnpmnnjvjnntmnmfftwwrfwrwswmmfrrfrrgbrbffwvvshvhrhmhththbbmqbmqqlslhssrmmqdmmjtmtmjtmjtttnwnvwvqwqjjnbbbdbqbnbpnbnllglcglcgcdczdznnqhhfthtmtlldqlqrrmddrldlzdllvddjcddqfqbqsbqqnllwppqpqzzrbbdppzsppjdpdqpdqdfqfrrwbwrwwqcqcsqsvvpbvbbztzptzzpccdtdhdffvqvcvzzmwzwddjfdffplplqlvlmmmvggpmpvpddpbptplpvlplvpvvnrvnnbqqqjhhwfhwfhwhqhmmpphqpqvppfzpzjzddgzzwffjmjggwhwwnnmlmpmmhcmcpcrcddvzvpzzwnznfznffgdgvddvtvgvsvdsdbbjnjtntbnttgbbbvgvgrgrzrvzrzddlsddcndcnnfqnnmpmlppdlplzplzpzgzmzmddlvlnnbttbwwhbhdhfdfssjppmcpplpdddnpdnnljlwjljsjnjhnhvhvqqsffrbbdttjdjndjdwwsfspffnhfhhlvhvmmqjmmwzwszwwvdvpdvdbdtdsdtsshvvmtvmtmctclchccrllznzfffpjjvhhdmhhvphpghgsgmmhlhnlnmnlnslnlgngznnsqnqddllpwllmzmjmttptfpplglqlgglgqqptqqmvmtmjmddcchbblltslsvsmvmgghmmccnzcztczzmnmttrdrvvcvzvvzllbhllnldndbbqffbbgtgddbtdbbzttdptpccjnjppbllbzlblrlcllhrrhqhgqqbcqcvcdvvnnzfzvzttrptrrwmrmlrlddvttdbtdbdcdvccwlcwwhphmppwfppclpcllgqgnghhvfflfggrzrcchfhhrdhrdhdnhnmmhjjwqjjpmmwvmmdnmnzzqmqwwmtthtdhtthnnqhqdhqqndqqwffsbspbptpmmndnllsmmdhmhfhnhjhghshlslppbgpgngddlsljsjmmzqzhhswhssfzssfqqcmqcmqcmmqggjcjvvgssrccwddmpmwwdfdpdbdpdwdvvqfvfrvvvsbvbhvvmqqcjqqvzqzppncnhhqnnpgplpqqpjpbblpbbbshsthhvfhfmmqzmmznnvrvqrrwdrdlrlwlttzqttjvttqltqqnfqqqwjqwwqttfstftjffsqqnhhnsnqqhggbsgghfgglslmssqlqhlhthqhccdsspsnssshbbnmngnnhllwclcffqllsrszrssnqsqvqjvjcvcttqgqbqmmfqfsqfsqswwvcvffndnfdfvfcvvggsmsfmfwfpfwwzhznntgtlglmmlfmllwrlrwwhcchqchhznzjjcdjdbjjhcjcscwwlnnsgngqqtgqgngnwgnnhqnnhchmchhtchcnclcmccgffbmmzvvrnngwwvddzccnjntjtwjwwztwtmtddjddpsptpbpbvvbwwnlnmndmnmdnnclnnbsbddpfdfvvjtjqqtqqqzjzlzqllzzwwlppvfffpcffffprrncnnzsnznhhwvvqhqphpjjgqqvnnmdmqqglqlblgglrlsspscsjjpvpbpjjwccslsppdjpdjjwvjjmhjhtjjwqqbqjqzjqzqpqbbswwlssqzssbjjpjqjbbjcjpjspjssjjzhhhnjhnhbnhhwzhzwlcshqlqpzgggzmcwntcwmfgtrwwjdpnbdqqcgnzgbdrzdmpwgvtvqffqbpvjpjrcfswffllnvnwvhclpjcwqwgnwqwvwsfgflrgzzsswffwjdjgvdvlgmczcbthwbvhggwzwlzfmhvwvjpbpnhcczbgfhhgghsmjwnvnsvnvmqwstrgnncwbqgbqpgdngllcqnzgwswpgtwzgqzggnzsdgltrlqfctqfvlzdswccfpdtjbfnrbqsmpjclnplbmqbmvwbzzdflwbqrljvzjpcrmnqsmrpqlmfsgcmthqpwgwzvmrjnhqczljcpnzjbwzrhjrzmcqpmlbzhgmqrlzsjbjsvcmcngptzlrthwsrjrlmsrgjlzrvpzwmprwnpgvjtspsppfvwfwcvbnqcwwmzlbqthqmbnbmnsnzgsbbnqtrvhlzjhphclpjzrdblszrnftqgwwrhpznhjhgrncvsvrmtmmgssvzddjfrnrzhbrqrfffjvzrqdnrdbvjwgrvlcpbncfgczlwdggsjmwzhndcdbggjvwfljctjnsjwczwfdrfttbhnlswfdbpcnwpspdhnzwqbgdswwpccbpfpgmfmvvwpzbzqsbbjbfnhjpszcbnrdplmwtdjtpcsztdjcmczltnstzwlcdbtdhsdgsgtlvdfqggfmmrppjfrmtfwhpbjsppszjbhmthndqmvbmqcbtqsltwrcvlvblwspbgspjftwllzcmnsrvjpnstzrfmcflnhppsdfggwbzvnvlnjqlfvrlplnzvfrwvgcgqvnpfgtgchctvhcplclzmfpwgnfhqjgglfmsgpflqcpqmbbhwnvvdllcnhblpnndbdtmgvfbvvvlvzlrpfqmnvzbfrssjtlgcjtpfznshvdjrjnfshfcgvwcdbqlfsbhnzwmsgwhpbzttgfjsqgwvdmbdwjljhsndrbbzfrsqjhcbldzqpmtnfvnmzltjcrvrltwshnhqlnclmcnfpbzstsczlqmfmdftzfbcwqnhqppzfbzpbfjhmmtvtbmblmtshsbtjlvsqvmbmgstbbdmhprqmtpfdqqntmnlbmpsmwfgrvstjcllhwpcddnljdjvdrbwqmgrjnldpgnrhgqpzqrvwzsngrgnbpjnsffzjsbdptwnnfcqlscfhvggpfstsnqzcjbqqhgdpqsrlprcppgqmddpqpbnvgwtdqsbbgtvsqfrtqfsbdzhsztfmvwrrsjcbtcjgzrnhnpgldtwbwgmwbgmjjzsbbzlhgmlczrhjwtzrgwscmjvlstprldhglvftqzbtrmcwzgtjppbnjcdvjvcwvdbngnbrmjvvtnwdqfclbpgzcfnnnlnngtgmhsqsdmbjctjzjpbrwrhscqshmmwbtfnzjgsrjlnqqdsvdrjdzsdprphnfmwwcztqfcrjvnfhlvnqwbrfmcvhrbtgvcrqjjfcnzwmlfzzdcbbzvphhmsdltwjfdcgthpvszqzjdbfwrpvhbjqdhrscnvjhjvvcldnhgjclmzpbrrwnscgpcqrpdgsnjnwhctcdqgwqbrcszfzpmtdrhlftvwffdjrtznqrppqbdbwvzmtlpvsqqpcngjgfdrpngnspdwhhvlhqrtsphgqrlldggtrvqsprbfdmrpgcmqphdvjfmhlznpgtqlvtnllcdhzhhtjjlfvdlwhcrfmjmdjtmbllvsfgvmfqtqlmrlrjmqptszvjdpzhphppljnpjdjpwlrclssgdnstchhwhpflmlrtdqvqbbljrmnflrltzpqmgqfrczvfzrpfsrwsgpljvjfjdjdvjchcdmmtjgghqspwzdtwqgtvmnrrbfbgnhcrvnzznrdlqmgmdwmpwzlqdjtvpszwnjtjtmjqvfwvftlhgpvgzswpbvbllfcwpjnsmbhzrdpdzjsrpnhphdcqjmzvvhrjcwhgwjwcshqwzpbpmfnjjvqcjrqmvsrdrtdvfhwhrbpvrqrsfzflslqtdrtcsggtzmpvbszdgttlvpwwltvpcwqmnwqtpcfzgsvsmncvpqqdrljfwtncplmjlpfcnqmcctwzhrbmrfwvsrjsbnhjrjmrnbmmnnhsvlltwzzhsgwppnlmljgpcsmpchdjdzpgvrtwsfzffhnlbfmrldzbshvpqhnfzpwnvczgfvhbntcpztwqlfgtsmdhvcrgjhvqrhbpvbpzcpbgzrcfjztbnfjptbzfpztwprwf",
	}

	for _, message := range messages {
		answer := -1

		for i := length; i <= len(message); i++ {
			substr := message[i-length : i]
			alldiff := true

			for i, ci := range substr {
				for j, cj := range substr {
					if i != j && ci == cj {
						alldiff = false
						break
					}
				}
			}

			if alldiff {
				answer = i
				break
			}

			//fmt.Println("substr", i, substr, "alldiff", alldiff)
		}

		fmt.Println("message", message, "answer", answer)
	}
}
