use std::fs::File;
use std::i32;
use std::io::{prelude::*, BufReader};

pub fn _part1(filename: &str) -> i32 {
    let races = _read_race_input_1(filename);
    
    let mut win_margin: usize = 1;

    for (race_time, winner_distance) in races {
        let mut wins: usize = 0;

        for secs_holding_button in 1..race_time {
            let distance = _calculate_traveled_distance(secs_holding_button, race_time);
            if distance > winner_distance {
                wins += 1;
            }
        }

        win_margin *= wins;
    }


    return win_margin as i32;
}

// d = distance
// s = speed && seconds pressing the button
// T = total race time
// d = s * (T - s)
fn _calculate_traveled_distance(secs_holding_button: usize, race_time: usize) -> usize {
    return secs_holding_button * (race_time - secs_holding_button);
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

fn _read_race_input_1(filename: &str) -> Vec<(usize, usize)> {
    let file = File::open(filename).unwrap();
    let mut reader: BufReader<File> = BufReader::new(file);

    let mut times: Vec<usize> = Vec::new();
    let mut distances: Vec<usize> = Vec::new();

    let mut buffer = String::new();
    let _ = reader.read_line(&mut buffer);
    
    let times_split = buffer.split_once(":").unwrap().1.split(" ").collect::<Vec<&str>>();
    for item in times_split {
        if item.trim().parse::<usize>().is_ok() {
            times.push(item.trim().parse::<usize>().unwrap());
        }
    }

    let mut buffer = String::new();
    let _ = reader.read_line(&mut buffer);
    
    let distances_split = buffer.split_once(":").unwrap().1.split(" ").collect::<Vec<&str>>();
    for item in distances_split {
        if item.trim().parse::<usize>().is_ok() {
            distances.push(item.trim().parse::<usize>().unwrap());
        }
    }

    let mut races: Vec<(usize, usize)> = Vec::new();
    for i in 0..times.len() {
        races.push((times[i], distances[i]));
    }

    return races;
}

fn _read_race_input_2(filename: &str) -> (usize, usize) {
    let file = File::open(filename).unwrap();
    let mut reader: BufReader<File> = BufReader::new(file);

    let mut buffer = String::new();
    let _ = reader.read_line(&mut buffer);
    
    let time_str = buffer.split_once(":").unwrap().1.replace(" ", "");
    let time = time_str.trim().parse::<usize>().unwrap();

    let mut buffer = String::new();
    let _ = reader.read_line(&mut buffer);
    
    let distance_str = buffer.split_once(":").unwrap().1.replace(" ", "");
    let distance = distance_str.trim().parse::<usize>().unwrap();

    return (time, distance);
}