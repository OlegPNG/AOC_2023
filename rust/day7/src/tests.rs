#[cfg(test)]
mod tests {
    use crate::{parse_file, PokerHand, HandType, get_hand_type};

    #[test]
    fn test_parse_file() {
        let data = parse_file("sample.txt").unwrap();
        let test_data = vec![
            PokerHand{
                hand: vec!['3', '2', 'T', '3', 'K'],
                bid: 765
            },
            PokerHand{
                hand: vec!['T', '5', '5', 'J', '5'],
                bid: 684
            },
            PokerHand{
                hand: vec!['K', 'K', '6', '7', '7'],
                bid: 28
            },
            PokerHand{
                hand: vec!['K', 'T', 'J', 'J', 'T'],
                bid: 220
            },
            PokerHand{
                hand: vec!['Q', 'Q', 'Q', 'J', 'A'],
                bid: 483
            },
        ];
        for i in 0 .. data.len() {
            assert_eq!(compare_hands(&test_data[i].hand, &data[i].hand), true);
            assert_eq!(&test_data[i].bid, &data[i].bid);
        }
    }
    fn compare_hands(expected: &Vec<char>, result: &Vec<char>) -> bool {
        for i in 0 .. result.len() {
            if expected[i] != result[i] {
                return false
            }
        }
        true
    }

    #[test]
    fn test_get_hand_type() {
        struct Test {
            hand: Vec<char>,
            expect: HandType
        }
        let tests = vec![
            Test{
                hand: vec!['3', '2', 'T', '3', 'K'],
                expect: HandType::OnePair
            },
            Test{
                hand: vec!['T', '5', '5', 'J', '5'],
                expect: HandType::ThreeKind
            },
            Test{
                hand: vec!['K', 'K', '6', '7', '7'],
                expect: HandType::TwoPair
            },
            Test{
                hand: vec!['K', 'T', 'J', 'J', 'T'],
                expect: HandType::TwoPair
            },
            Test{
                hand: vec!['Q', 'Q', 'Q', 'J', 'A'],
                expect: HandType::ThreeKind
            }
        ];

        for test in tests {
            let result = get_hand_type(&test.hand);
            assert_eq!(test.expect, result);
        }
    }
}
