
mod challenge1;
mod challenge2;
mod challenge3;
mod challenge4;
mod challenge5;
mod challenge6;
mod challenge7;
mod common;

fn main() {
    let challenge_number = 7;
    let input_type = "sample";
    let filename = format!("../res/challenge{}/{}.txt", challenge_number, input_type);
    
    // use std::time::Instant;
    // let now = Instant::now();
    // let result = challenge7::_part1(&filename);
    // println!("Part 1 result: {}", result);
    // let elapsed = now.elapsed();
    // println!("Part 1 elapsed time: {:?}", elapsed);

    // let now = Instant::now();
    let result = challenge7::_part2(&filename);
    println!("Part 2 result: {}", result);
    // let elapsed = now.elapsed();
    // println!("Part 2 elapsed time: {:?}", elapsed);
}