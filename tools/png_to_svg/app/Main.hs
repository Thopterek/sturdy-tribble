import Codec.Picture
import System.FilePath ((</>))

-- Bro what is this as for saving the path we need extra import hahah
path :: FilePath
path = "../../overview/pixel_art_sketches/"

-- Haskell, why are you using tabs instead of spaces :c
main :: IO ()
main = do
    result <- readImage (path </> "bird.png")
    case result of
        Left err -> putStrLn $ "Error: " ++ err
        Right img -> do
            let w = dynamicMap imageWidth img
            let h = dynamicMap imageHeight img
            putStrLn $ "Width: " ++ show w ++ ", Height: " ++ show h
