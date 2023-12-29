mod tests;
use std::collections::HashMap;
use std::{fs::File, io::BufReader};
use std::io::prelude::*;

#[derive(Clone)]
pub struct PokerHand {
    pub hand: Vec<char>,
    pub bid: u32
}

//#[repr(u8)]
#[derive(Debug, PartialEq, Eq, PartialOrd, Clone)]
pub enum HandType {
    HighCard = 0,
    OnePair = 1,
    TwoPair = 2,
    ThreeKind = 3,
    FullHouse = 4,
    FourKind = 5,
    FiveKind = 6
}

impl From<u8> for HandType {
    fn from(value: u8) -> Self {
        match value {
            1 => Self::OnePair,
            2 => Self::TwoPair,
            3 => Self::ThreeKind,
            4 => Self::FullHouse,
            5 => Self::FourKind,
            6 => Self::FiveKind,
            _ => Self::HighCard
        }
    }
}

/*impl fmt::Display for HandType {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match *self {
            HandType::HighCard => ,
            HandType::OnePair => todo!(),
            HandType::TwoPair => todo!(),
            HandType::ThreeKind => todo!(),
            HandType::FullHouse => todo!(),
            HandType::FourKind => todo!(),
            HandType::FiveKind => todo!(),
        }
    }
}*/

#[derive(Clone)]
pub struct HandData {
    pub poker_hand: PokerHand,
    pub hand_type: HandType
}

pub fn parse_file(file_name: &str ) -> Result<Vec<PokerHand>, Box<dyn std::error::Error>> {
    let file = File::open(file_name)?;
    let reader = BufReader::new(file);

    let mut hands: Vec<PokerHand> = Vec::new();

    for line_result in reader.lines() {
        let line = line_result?;
        //println!("{}", line);
        let split_string: Vec<&str> = line.split(" ").collect();
        let hand = split_string[0].chars().collect();
        let bid: u32 = split_string[1].parse()?; 
        hands.push(PokerHand{
            hand, bid
        });
    }
    //println!("Total Hands: {}", hands.len());
    Ok(hands)
}

pub fn get_card_value(card: &char) -> u8 {
    let card_values = HashMap::from([
        ('J', 0),
        ('2', 1),
        ('3', 2),
        ('4', 3),
        ('5', 4),
        ('6', 5),
        ('7', 6),
        ('8', 7),
        ('9', 8),
        ('T', 9),
        ('Q', 10),
        ('K', 11),
        ('A', 12),
    ]);
    return *card_values.get(card).unwrap();
}

pub fn get_hand_type(hand: &Vec<char>) -> HandType {
    let mut card_map: HashMap<char, u32> = HashMap::new();
    for card in hand {
        card_map.entry(*card).and_modify(|counter| *counter += 1).or_insert(1);
    }

    let mut three_kind_present = false;
    let mut one_pair_present = false;
    let mut hand_type = HandType::HighCard;
    for (_, value) in card_map {
        match value {
            2 => { 
                if one_pair_present == true {
                    if HandType::TwoPair > hand_type {
                        hand_type = HandType::TwoPair;
                        continue;
                    }
                }
                if three_kind_present {
                    if HandType::FullHouse > hand_type {
                        hand_type = HandType::FullHouse;
                        continue;
                    }
                }
                one_pair_present = true;
                if HandType::OnePair > hand_type {
                    hand_type = HandType::OnePair;
                }
            },
            3 => {
                if one_pair_present {
                    if HandType::FullHouse > hand_type {
                        hand_type = HandType::FullHouse;
                        continue
                    }
                }
                three_kind_present = true;
                if HandType::ThreeKind > hand_type {
                    hand_type = HandType::ThreeKind;
                }
            },
            4 => {
                if HandType::FourKind > hand_type {
                    hand_type = HandType::FourKind;
                }  
            },
            5 => {
                hand_type = HandType::FiveKind;
                return hand_type
            }
            _ => {}
        };
    };
    return hand_type
}

pub fn get_hand_type_using_joker(hand: &Vec<char>) -> HandType {
    let mut card_map: HashMap<char, u32> = HashMap::new();
    let mut card_totals: Vec<u32> = Vec::new();
    let mut joker_count = 0;
    for card in hand {
        card_map.entry(*card).and_modify(|counter| *counter += 1).or_insert(1);
        if *card == 'J' {
            joker_count += 1;
        }
    }
    if joker_count == 4 || joker_count == 5 {
        return HandType::FiveKind
    }

    let mut three_kind_present = false;
    let mut one_pair_present = false;
    let mut hand_type = HandType::HighCard;

    let max = card_map.clone().into_iter().max().unwrap();

    let mut card_count;
    for (key, value) in card_map {
        if key != 'J' {
            card_count = value + joker_count;
        } else {
            card_count = value;
        }
        match card_count {
            2 => {
                if key != max.0 {
                    
                }
            }
            _ => {}
        }
    }

    return hand_type

}

fn compare_hands(left: HandData, right: HandData) -> (HandData, HandData) {
    if left.hand_type > right.hand_type {
        // Swap left and right
        return (right, left)
    } else if left.hand_type < right.hand_type {
        // Return original order
        return (left, right)
    } else {
        let left_hand = &left.poker_hand.hand;
        let right_hand = &right.poker_hand.hand;
        for i in 0..left_hand.len() {
            let left_value = get_card_value(&left_hand[i]);
            let right_value = get_card_value(&right_hand[i]);
            if left_value == right_value {
                continue;
            }
            if left_value > right_value {
                return (right, left)
            }
            return (left, right)
        }
    }

    return (left, right)
}
fn bubble_sort(hands: &mut Vec<HandData>) {
    for i in (0..hands.len()).rev() {
        for j in 0..i {
            (hands[j], hands[j+1]) = compare_hands(hands[j].clone(), hands[j+1].clone());
        }
    }
}

fn part_one(data: Vec<PokerHand>) {
    let mut hand_data: Vec<HandData> = Vec::new();
    for hand in data {
        let hand_type = get_hand_type(&hand.hand);
        hand_data.push(
            HandData { poker_hand: hand, hand_type }
        );
    }
    bubble_sort(&mut hand_data);
    let mut total = 0;
    for i in 0..hand_data.len() {
        let value = (i as u32 + 1) * hand_data[i].poker_hand.bid;
        let hand: String = hand_data[i].poker_hand.hand.iter().collect();
        println!("Hand {}: {}", i, hand);
        total += value;
    }
    println!("Total: {}", total);
}

fn part_two(data: Vec<PokerHand>) {

}

fn main() {
    let hands = parse_file("input.txt").unwrap();
    part_one(hands);
}
