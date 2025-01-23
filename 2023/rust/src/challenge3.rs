
    use std::collections::HashMap;
    use std::i32;

    use crate::common::_read_map_matrix;

    pub fn _challenge3_1(filename: &str) -> i32 {
        let matrix_map: Vec<Vec<char>> = _read_map_matrix(filename);
    
        let mut digits_indexes_map : HashMap<Vec<(i32,i32)>, usize> = HashMap::new();
        let mut symbols_indexes_map : HashMap<(i32,i32), bool> = HashMap::new();
    
        for (i, line) in matrix_map.iter().enumerate() {
            let mut j: usize = 0;
    
            while j < line.len() {
                if !line[j].is_digit(10) && line[j] != '.' {
                    symbols_indexes_map.insert((i as i32, j as i32), true);
                    j+=1;
                } else if line[j].is_digit(10) {
                    let mut digits: String = String::new();
                    let mut coordinates: Vec<(i32,i32)> = Vec::new();
                    while j < line.len() && line[j].is_digit(10) {
                        digits.push(line[j]);
                        coordinates.push((i as i32, j as i32));
                        j +=1;
                    }
    
                    let number = digits.parse::<usize>().unwrap();
                    digits_indexes_map.insert(coordinates, number);               
                } else {
                    j+=1;
                }
            }
        }
    
        let mut final_sum: i32 = 0;
    
        for (coordinates, value) in digits_indexes_map.iter() {
            let mut belongs = false;
            
            for (x,y) in coordinates {
                if symbols_indexes_map.contains_key(&(*x, *y+1)) ||
                symbols_indexes_map.contains_key(&(*x, *y-1)) ||
                symbols_indexes_map.contains_key(&(*x+1, *y+1)) ||
                symbols_indexes_map.contains_key(&(*x+1, *y)) ||
                symbols_indexes_map.contains_key(&(*x+1, *y-1)) ||
                symbols_indexes_map.contains_key(&(*x-1, *y+1)) ||
                symbols_indexes_map.contains_key(&(*x-1, *y)) ||
                symbols_indexes_map.contains_key(&(*x-1, *y-1)){
                    belongs = true;
                    break;
                }
            }
    
            if belongs {
                final_sum += *value as i32;
            }
        }
    
        return final_sum;
    }
    
    pub fn _challenge3_2(filename: &str) -> i32 {
        let matrix_map: Vec<Vec<char>> = _read_map_matrix(filename);
    
        let mut digits_indexes_map : HashMap<Vec<(i32,i32)>, usize> = HashMap::new();
        let mut gears_indexes_map : HashMap<(i32,i32), i32> = HashMap::new();
        let mut gears_and_numbers : HashMap<(i32,i32), Vec<i32>> = HashMap::new();
    
        for (i, line) in matrix_map.iter().enumerate() {
            let mut j: usize = 0;
    
            while j < line.len() {
                if line[j] == '*' {
                    gears_indexes_map.insert((i as i32, j as i32), 0);
                    j+=1;
                } else if line[j].is_digit(10) {
                    let mut digits: String = String::new();
                    let mut coordinates: Vec<(i32,i32)> = Vec::new();
                    while j < line.len() && line[j].is_digit(10) {
                        digits.push(line[j]);
                        coordinates.push((i as i32, j as i32));
                        j +=1;
                    }
    
                    let number = digits.parse::<usize>().unwrap();
                    digits_indexes_map.insert(coordinates, number);               
                } else {
                    j+=1;
                }
            }
        }
    
        for (coordinates, number) in digits_indexes_map.iter() {
            
            for (x,y) in coordinates {
                if gears_indexes_map.contains_key(&(*x, *y+1)) {
                    match gears_and_numbers.get_mut(&(*x, *y+1)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x, *y+1), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x, *y-1)) {
                    match gears_and_numbers.get_mut(&(*x, *y-1)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x, *y-1), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x+1, *y+1)) {
                    match gears_and_numbers.get_mut(&(*x+1, *y+1)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x+1, *y+1), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x+1, *y)) {
                    match gears_and_numbers.get_mut(&(*x+1, *y)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x+1, *y), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x+1, *y-1)) {
                    match gears_and_numbers.get_mut(&(*x+1, *y-1)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x+1, *y-1), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x-1, *y+1)) {
                    match gears_and_numbers.get_mut(&(*x-1, *y+1)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x-1, *y+1), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x-1, *y)) {
                    match gears_and_numbers.get_mut(&(*x-1, *y)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x-1, *y), vec![*number as i32]);
                        }
                    }
                    break;
                } else if gears_indexes_map.contains_key(&(*x-1, *y-1)){
                    match gears_and_numbers.get_mut(&(*x-1, *y-1)) {
                        Some(numbers) => {
                            numbers.push(*number as i32);
                        },
                        None => {
                            gears_and_numbers.insert((*x-1, *y-1), vec![*number as i32]);
                        }
                    }
                    break;
                }
            }
        }
    
        let mut final_sum: i32 = 0;
        for (_, numbers) in gears_and_numbers.iter() {
            if numbers.len() == 2 {
                final_sum += numbers[0] * numbers[1];
            }
        }
    
        return final_sum;
    }