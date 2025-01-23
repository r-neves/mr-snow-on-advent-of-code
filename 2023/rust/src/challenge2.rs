
    use std::collections::HashMap;
    use std::fs::File;
    use std::i32;
    use std::io::{prelude::*, BufReader};

    pub fn _challenge2_1(filename: &str) -> i32 {
        let max_values = HashMap::from([("red", 12), ("green", 13), ("blue", 14)]);
    
        let file = File::open(filename).unwrap();
        let reader: BufReader<File> = BufReader::new(file);
    
        let mut possible_games_sum: i32 = 0;
    
        for line in reader.lines() {
            let mut is_game_possible = true;
            let line = line.unwrap();
            
            let (game_prefix, game) = line.split_once(":").unwrap();
            let game_number = game_prefix.split_once(" ").unwrap().1;
    
            for (_, round) in game.split(";").into_iter().enumerate() {
                for (_, play) in round.split(",").into_iter().enumerate() {
                    let trimmed_play = play.trim();
                    let (num_cubes, color) = trimmed_play.split_once(" ").unwrap();
                    let i32_num_cubes = num_cubes.parse::<i32>().unwrap();
    
                    if max_values.get(color).unwrap() < &i32_num_cubes {
                        is_game_possible = false;
                        break;
                    }
                }
    
                if !is_game_possible {
                    break;
                }
            }
    
            if is_game_possible {
                possible_games_sum += game_number.parse::<i32>().unwrap();
            }
        }
    
        possible_games_sum
    }
    
    pub fn _challenge2_2(filename: &str) -> i32 {
        let file = File::open(filename).unwrap();
        let reader: BufReader<File> = BufReader::new(file);
    
        let mut games_power_sum: i32 = 0;
    
        for line in reader.lines() {
            let line = line.unwrap();
            
            let (_, game) = line.split_once(":").unwrap();
    
            let mut max_cubes = HashMap::from([("red", 0), ("green", 0), ("blue", 0)]);
    
            for (_, round) in game.split(";").into_iter().enumerate() {
                for (_, play) in round.split(",").into_iter().enumerate() {
                    let trimmed_play = play.trim();
                    let (num_cubes, color) = trimmed_play.split_once(" ").unwrap();
                    let i32_num_cubes = num_cubes.parse::<i32>().unwrap();
    
                    if max_cubes.get(color).unwrap() < &i32_num_cubes {
                        max_cubes.insert(color, i32_num_cubes);
                    }
                }
            }
    
            let mut game_product_power: i32 = 1;
            for (_, (_, value)) in max_cubes.iter().enumerate() {
                game_product_power *= value;
            }
    
            games_power_sum += game_product_power;
        }
    
        games_power_sum
    }