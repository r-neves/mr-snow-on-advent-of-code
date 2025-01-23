use std::collections::HashMap;
use std::fs::File;
use std::i32;
use std::io::{prelude::*, BufReader};

pub fn _challenge4_1(filename: &str) -> i32 {
    let file = File::open(filename).unwrap();
    let reader: BufReader<File> = BufReader::new(file);

    let mut sum: usize = 0;

    for line in reader.lines() {
        let line = line.unwrap();
        let (win_numbers_map, card_numbers) = _parse_card_line(line);

        let mut points: usize = 0;
        for number in card_numbers {
            if win_numbers_map.contains_key(&number) {
                if points == 0 {
                    points = 1;
                } else {
                    points *= 2;
                }
            }
        }

        sum += points;
    }

    return sum as i32;
}

fn _parse_card_line(line: String) -> (HashMap<usize, bool>, Vec<usize>) {
    let mut winning_numbers_map: HashMap<usize, bool> = HashMap::new();
    let mut card_numbers_vec: Vec<usize> = Vec::new();

    let card = line.split_once(":").unwrap().1;
    let (winning_numbers_str, card_numbers_str) = card.split_once("|").unwrap();
    
    for number in winning_numbers_str.split(" ").enumerate() {
        if number.1.is_empty() {
            continue;
        }

        let number = number.1.parse::<usize>().unwrap();
        winning_numbers_map.insert(number, true);
    }

    for number in card_numbers_str.split(" ").enumerate() {
        if number.1.is_empty() {
            continue;
        }
        
        let number = number.1.parse::<usize>().unwrap();
        card_numbers_vec.push(number);
    }

    return (winning_numbers_map, card_numbers_vec);
}

pub fn _challenge4_2(filename: &str) -> i32 {
    let file = File::open(filename).unwrap();
    let reader: BufReader<File> = BufReader::new(file);

    let mut scratch_card_map: HashMap<usize, usize> = HashMap::new();
    let mut line_number: usize = 1;

    for line in reader.lines() {
        
        match scratch_card_map.get_key_value(&line_number) {
            Some((_, v)) => { scratch_card_map.insert(line_number, v + 1); },
            None => { scratch_card_map.insert(line_number, 1); },
        }

        let line = line.unwrap();
        let (win_numbers_map, card_numbers) = _parse_card_line(line);

        let mut points: usize = 0;
        for number in card_numbers {
            if win_numbers_map.contains_key(&number) {
                points += 1;
            }
        }

        let multiplier = scratch_card_map.get_key_value(&line_number).unwrap().1.clone(); 
        if points > 0 {     
            for i in 1..=points {
                let key = line_number + i;
                match scratch_card_map.get_key_value(&key) {
                    Some((_,v)) => { scratch_card_map.insert(key, v + multiplier); },
                    None => { scratch_card_map.insert(key, multiplier); },
                }
            }
        }
        
        line_number += 1;
    }

    let mut sum: usize = 0;
    for (_, value) in scratch_card_map.iter() {
        sum += value;
    }

    return sum as i32;
}