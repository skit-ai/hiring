#!/usr/bin/env fish

for i in (seq 0 4)
    for fmt in mp3 aac ogg
        ffmpeg -i $i.wav $i.$fmt
        ffmpeg -i $i.$fmt $i.$fmt.wav
        rm $i.$fmt
    end
end
