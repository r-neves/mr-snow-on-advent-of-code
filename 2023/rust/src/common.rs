use std::fs::File;
use std::io::*;
use std::path::Path;

pub fn _read_map_matrix(filename: &str) -> Vec<Vec<char>> {
        let file = File::open(filename).unwrap();
        let reader: BufReader<File> = BufReader::new(file);
    
        let mut map_matrix: Vec<Vec<char>> = Vec::new();
    
        for line_opt in reader.lines() {
            let line = line_opt.unwrap();
            let mut line_vec: Vec<char> = Vec::new();
    
            for c in line.chars() {
                line_vec.push(c);
            }
    
            map_matrix.push(line_vec);
        }
    
        map_matrix
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn _read_lines<P>(filename: P) -> Result<Lines<BufReader<File>>> where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(BufReader::new(file).lines())
}