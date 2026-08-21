use image::{DynamicImage, GenericImageView};

use rust_png_to_svg::c_to_char;
#[cfg(debug_assertions)]
use rust_png_to_svg::print_png_on_terminal;

fn main() {
    let path_to_img = "../../overview/pixel_art_sketches/bird.png";
    let res = image::open(path_to_img);
    if let Ok(img) = res {
        let dimensions = img.dimensions();
        println!("Dimensions {:?}", dimensions);
        if dimensions.0 != 16 && dimensions.1 != 16 {
            println!("Error: Wrong image my guy, here we work on 16 by 16");
            println!("Info: Should it be an error? No. Do I care? Not enough");
            return;
        }
        println!("horizontal rec size {}", horizontal_rectangle(&img));
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

fn horizontal_rectangle(img_slice: &DynamicImage) -> usize {
    let mut size: usize = 1;
    let mut offset: usize = 0;
    let mut first_pixel: char = 'N';
    for (counter, (_w, _h, pixel_data)) in img_slice.pixels().enumerate() {
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

/*
fn vertical_rectangle(img: &DynamicImage) -> (usize) {
    0
}
*/

/*
* Actually change of idea let's to do it easier
fn find_rectangle(img_slice: &DynamicImage) -> (usize, usize) {
    let mut end_x: usize = 0;
    let mut end_y: usize = 0;
    let mut n_pixel: String;
    let mut last_p: String;
    let mut pixel_changed = false;
    for (counter, pixel_data) in img_slice.pixels().enumerate() {
        if pixel_data.2[4] >= 127 {
            n_pixel = get_symbol_for_terminal(pixel_data.2[0]);
            if n_pixel != last_p {
                pixel_changed = true;
            }
        }
        end_x = counter - (end_y * 15);
        if (counter + 1) % 16 == 0 {
            end_y += 1;
            #[cfg(debug_assertions)]
            print!("next_line");
        }
    }
    (end_x, end_y)
}
*/
