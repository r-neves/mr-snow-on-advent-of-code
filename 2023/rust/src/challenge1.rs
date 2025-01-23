    use std::collections::HashMap;
    use std::fs::File;
    use std::i32;
    use std::io::{prelude::*, BufReader};

    pub fn _challenge1_1(filename: &str) -> i32 {
        let file = File::open(filename).unwrap();
        let reader: BufReader<File> = BufReader::new(file);
    
        let mut sum: i32 = 0;
    
        for line in reader.lines() {
            let line = line.unwrap();
            
            let mut first_digit: i32 = -1;
            let mut last_digit: i32 = -1;
    
            for (_, c) in line.chars().enumerate() {
                if c.is_numeric() {
                    if first_digit == -1 {
                        first_digit = c.to_digit(10).unwrap() as i32;
                    }
                    last_digit = c.to_digit(10).unwrap() as i32;
                }
            }
    
            first_digit = first_digit * 10 + last_digit;
            sum += first_digit;
        }
    
        sum
    }
    
    pub fn _challenge1_2(filename: &str) -> i32 {
        let file = File::open(filename).unwrap();
        let reader: BufReader<File> = BufReader::new(file);
    
        let mut sum: i32 = 0;
        let number_words = HashMap::from([("one", 1), ("two", 2), ("three", 3), ("four", 4), ("five", 5), ("six", 6), ("seven", 7), ("eight", 8), ("nine", 9)]);
    
        for line in reader.lines() {
            let line = line.unwrap();
            
            let mut first_digit: i32 = 0;
            let mut first_digit_index: i32 = line.len() as i32;
            let mut last_digit: i32 = 0;
            let mut last_digit_index: i32 = 0;
    
            for (i, c) in line.chars().enumerate() {
                if let Some(digit) = c.to_digit(10) {
                        if (i as i32) < first_digit_index {
                            first_digit = digit as i32;
                            first_digit_index = i as i32;
                        }
        
                        if (i as i32) > last_digit_index {
                            last_digit = digit as i32;
                            last_digit_index = i as i32;
                        }
                }
            }
    
            for (_, (key, value)) in number_words.iter().enumerate() {
                match line.find(key) {
                    Some(index) => {
                        let idx = index as i32;
                        if idx < first_digit_index {
                            first_digit = *value;
                            first_digit_index = idx;
                        } else {}
                    },
                    None => {}
                }
    
                match line.rfind(key) {
                    Some(index) => {
                        let idx = index as i32;
                        if idx > last_digit_index {
                            last_digit = *value;
                            last_digit_index = idx;
                        }
                    },
                    None => {}
                }
            }
    
            first_digit = first_digit * 10 + last_digit;
            sum += first_digit;
        }
        
        sum
    }
    