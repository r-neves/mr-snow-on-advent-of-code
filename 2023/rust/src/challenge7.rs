use std::fs::File;
use std::i32;
use std::io::{prelude::*, BufReader};

const CARDS: [&str; 13] = ["2","3","4","5","6","7","8","9","T","J","Q","K","A"];

pub fn _part1(filename: &str) -> i32 {
    let file = File::open(filename).unwrap();
    let mut reader: BufReader<File> = BufReader::new(file);

    let mut lines = Vec::new();

    for line in reader.lines() {
        lines.push(line.unwrap());
    }
    
    

    let final_sum: usize = 0;

    return final_sum as i32;
}

fn parse_hand_and_value(line: &str) -> Vec<(usize, usize)> {
    let mut result: Vec<(usize, usize)> = Vec::new();

    let (hand, value_str) = line.split_once(" ").unwrap();
    let value = value_str.parse::<usize>().unwrap();

    return result;
}

pub fn _part2(filename: &str) -> i32 {
    let (race_time, winner_distance) = _read_race_input_2(filename);

    let mut wins: usize = 0;

    for secs_holding_button in 1..race_time {
        let distance = _calculate_traveled_distance(secs_holding_button, race_time);
        if distance > winner_distance {
            wins += 1;
        }
    }

    return wins as i32;
}
