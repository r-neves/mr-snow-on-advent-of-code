use std::fs::File;
use std::i32;
use std::io::{prelude::*, BufReader};

#[derive(Debug)]
struct _GardeningInput1 {
    seeds: Vec<usize>,
    seed_to_soil_maps: Vec<MapValues>,
    soil_to_fertilizer_maps: Vec<MapValues>,
    fertilizer_to_water_maps: Vec<MapValues>,
    water_to_light_maps: Vec<MapValues>,
    light_to_temperature_maps: Vec<MapValues>,
    temperature_to_humidity_maps: Vec<MapValues>,
    humidity_to_location_maps: Vec<MapValues>,
}

struct _GardeningInput2 {
    seeds_pairs: Vec<(usize, usize)>,
    seed_to_soil_maps: Vec<MapValues>,
    soil_to_fertilizer_maps: Vec<MapValues>,
    fertilizer_to_water_maps: Vec<MapValues>,
    water_to_light_maps: Vec<MapValues>,
    light_to_temperature_maps: Vec<MapValues>,
    temperature_to_humidity_maps: Vec<MapValues>,
    humidity_to_location_maps: Vec<MapValues>,
}

#[derive(Debug)]
struct MapValues {
    source_start: usize,
    source_end: usize,
    destination_start: usize,
    destination_end: usize,
    range_length: usize,
}

pub fn _challenge5_1(filename: &str) -> i32 {
    let gardening_input = _parse_garden_input_1(filename);

    let mut min_location = usize::MAX;

    for seed in gardening_input.seeds {
        let soil = _get_translation_value(seed, &gardening_input.seed_to_soil_maps);
        let fertilizer = _get_translation_value(soil, &gardening_input.soil_to_fertilizer_maps);
        let water = _get_translation_value(fertilizer, &gardening_input.fertilizer_to_water_maps);
        let light = _get_translation_value(water, &gardening_input.water_to_light_maps);
        let temperature = _get_translation_value(light, &gardening_input.light_to_temperature_maps);
        let humidity = _get_translation_value(temperature, &gardening_input.temperature_to_humidity_maps);
        let location = _get_translation_value(humidity, &gardening_input.humidity_to_location_maps);

        if location < min_location {
            min_location = location;
        }
    }

    return min_location as i32;
}

fn _get_translation_value(source: usize, mappings: &Vec<MapValues>) -> usize {
    for mapping in mappings {
        if source >= mapping.source_start && source <= mapping.source_end {
            return mapping.destination_start + (source - mapping.source_start);
        }
    }

    return source;
}

fn _get_reversed_translation_value(destination: usize, mappings: &Vec<MapValues>) -> usize {
    for mapping in mappings {
        if destination >= mapping.destination_start && destination <= mapping.destination_end {
            let diff = destination - mapping.destination_start;
            return mapping.source_start + diff;
        }
    }

    return destination;
}

pub fn _challenge5_2(filename: &str) -> i32 {
    let gardening_input: _GardeningInput2 = _parse_garden_input_2(filename);

    let mut location_max = usize::MAX;
    for location_mapping in gardening_input.humidity_to_location_maps.iter() {
        if location_mapping.destination_end > location_max {
            location_max = location_mapping.destination_end;
        }
    }

    for location in 0..location_max {
        let humidity = _get_reversed_translation_value(location, &gardening_input.humidity_to_location_maps);
        let temperature = _get_reversed_translation_value(humidity, &gardening_input.temperature_to_humidity_maps);
        let light = _get_reversed_translation_value(temperature, &gardening_input.light_to_temperature_maps);
        let water = _get_reversed_translation_value(light, &gardening_input.water_to_light_maps);
        let fertilizer = _get_reversed_translation_value(water, &gardening_input.fertilizer_to_water_maps);
        let soil = _get_reversed_translation_value(fertilizer, &gardening_input.soil_to_fertilizer_maps);
        let seed = _get_reversed_translation_value(soil, &gardening_input.seed_to_soil_maps);

        for seed_pair in gardening_input.seeds_pairs.iter() {
            if seed >= seed_pair.0 && seed < seed_pair.0 + seed_pair.1 {
                return location as i32;
            }
        }
    }

    return -1 as i32;
}

fn _parse_garden_input_1(filename: &str) -> _GardeningInput1 {
    let file = File::open(filename).unwrap();
    let mut reader: BufReader<File> = BufReader::new(file);

    let mut seeds_line = String::new();
    let _ = reader.read_line(&mut seeds_line);
    seeds_line = seeds_line[..seeds_line.len()-1].to_string();

    let seeds_str = seeds_line.split_once(": ").unwrap().1.split(" ").collect::<Vec<&str>>();
    let seeds = seeds_str.iter().map(|s| s.parse::<usize>().unwrap()).collect::<Vec<usize>>();

    // Consume empty line
    let _ = reader.read_line(&mut String::new());

    return _GardeningInput1 {  
        seeds,
        seed_to_soil_maps: _read_mappings_group(&mut reader),
        soil_to_fertilizer_maps: _read_mappings_group(&mut reader),
        fertilizer_to_water_maps: _read_mappings_group(&mut reader),
        water_to_light_maps: _read_mappings_group(&mut reader),
        light_to_temperature_maps: _read_mappings_group(&mut reader),
        temperature_to_humidity_maps: _read_mappings_group(&mut reader),
        humidity_to_location_maps: _read_mappings_group(&mut reader),
    };
}

fn _read_mappings_group(reader: &mut BufReader<File>) -> Vec<MapValues> {
    let mut mapping_values = Vec::<MapValues>::new();

    // Consume header line
    let _ = reader.read_line(&mut String::new());

    let mut buffer = String::new();
    let _ = reader.read_line(&mut buffer);
    while buffer.len() > 1 {
        if buffer.ends_with("\n") {
            buffer = buffer[..buffer.len()-1].to_string();
        }
        let mapping_split: Vec<&str> = buffer.splitn(3, " ").collect();

        let dest_start = mapping_split[0].parse::<usize>().unwrap();
        let src_start = mapping_split[1].parse::<usize>().unwrap();
        let range = mapping_split[2].parse::<usize>().unwrap();
        let mapping = MapValues {
            source_start: src_start,
            source_end: src_start + range,
            destination_start: dest_start,
            destination_end: dest_start + range,
            range_length: range,
        };

        mapping_values.push(mapping);

        buffer.clear();
        let _ = reader.read_line(&mut buffer);
    }

    return mapping_values;
}

fn _parse_garden_input_2(filename: &str) -> _GardeningInput2 {
    let file = File::open(filename).unwrap();
    let mut reader: BufReader<File> = BufReader::new(file);

    let mut seeds_line = String::new();
    let _ = reader.read_line(&mut seeds_line);
    seeds_line = seeds_line[..seeds_line.len()-1].to_string();

    let seeds_str = seeds_line.split_once(": ").unwrap().1.split(" ").collect::<Vec<&str>>();
    let seeds = seeds_str.iter().map(|s| s.parse::<usize>().unwrap()).collect::<Vec<usize>>();

    let mut seeds_pairs = Vec::<(usize, usize)>::new();
    let mut i = 0;
    while i < seeds.len() {
        seeds_pairs.push((seeds[i], seeds[i+1]));
        i += 2;
    } 

    // Consume empty line
    let _ = reader.read_line(&mut String::new());

    let seed_to_soil_maps = _read_mappings_group(&mut reader);

    let soil_to_fertilizer_maps = _read_mappings_group(&mut reader);

    let fertilizer_to_water_maps = _read_mappings_group(&mut reader);

    let water_to_light_maps = _read_mappings_group(&mut reader);

    let light_to_temperature_maps = _read_mappings_group(&mut reader);

    let temperature_to_humidity_maps = _read_mappings_group(&mut reader);

    let humidity_to_location_maps = _read_mappings_group(&mut reader);

    return _GardeningInput2 {  
        seeds_pairs,
        seed_to_soil_maps,
        soil_to_fertilizer_maps,
        fertilizer_to_water_maps,
        water_to_light_maps,
        light_to_temperature_maps,
        temperature_to_humidity_maps,
        humidity_to_location_maps,
    }
}