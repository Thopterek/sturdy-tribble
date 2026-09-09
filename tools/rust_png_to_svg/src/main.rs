use std::usize;

use image::{DynamicImage, GenericImageView};

use rust_png_to_svg::c_to_char;
#[cfg(debug_assertions)]
use rust_png_to_svg::print_png_on_terminal;

fn double_check_failed(img: &DynamicImage) -> bool {
    let dimensions = img.dimensions();
    println!("Dimensions {:?}", dimensions);
    if dimensions.0 != 16 && dimensions.1 != 16 {
        println!("Error: Wrong image my guy, here we work on 16 by 16");
        println!("Info: Should it be an error? No. Do I care? Not enough");
        return true;
    }
    false
}

fn main() {
    let path_to_img = "../../overview/pixel_art_sketches/bird.png";
    let res = image::open(path_to_img);
    if let Ok(img) = res {
        if double_check_failed(&img) {
            return;
        }
        println!("horizontal rec size {}", horizontal_rectangle(&img, 0));
        println!("vertical rec size {}", vertical_rectangle(&img, 0));
        #[cfg(debug_assertions)]
        print_png_on_terminal(img);
        //print_rgba(img);
    } else {
        println!("Error: It uses relative path to find image {}", path_to_img);
    }
}

fn debug_vals(c: char, size: usize, offset: usize) {
    println!("{}, size {}, offset {}", c, size, offset);
}

/*
* I am not handling two lines of same color because I didn't draw any
* if it would be actually useful I should probably add it, just like many other things
*/
fn horizontal_rectangle(img: &DynamicImage, start: usize) -> usize {
    let mut size: usize = 1;
    let mut offset: usize = 0;
    let mut first_pixel: char = 'N';
    for (mut counter, (_w, _h, pixel_data)) in img.pixels().enumerate() {
        if counter == 0 {
            counter = start;
        }
        if pixel_data[3] >= 127 {
            size = counter + 1;
            debug_vals(c_to_char(pixel_data[0]), size, offset);
            if first_pixel == 'N' {
                first_pixel = c_to_char(pixel_data[0]);
                offset = counter
            }
            if c_to_char(pixel_data[0]) != first_pixel {
                break;
            }
        }
        if (counter + 1) % 16 == 0 {
            break;
        }
    }
    size - offset
}

fn vertical_rectangle(img: &DynamicImage, start: usize) -> usize {
    let mut size: usize = 0;
    let mut f_pixel: char = 'N';
    let mut current_p: char = 'N';
    while f_pixel == current_p {
        for (counter, (_w, _h, pixel_data)) in img.pixels().enumerate() {
            current_p = c_to_char(pixel_data[0]);
            if (f_pixel == 'N') {
                f_pixel = c_to_char(pixel_data[0]);
            }
        }
        size += 1;
    }
    size
}
