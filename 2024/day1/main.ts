import {readLines} from "https://deno.land/std/io/mod.ts";

async function solution1() {
    const file = await Deno.open("./input.txt");

    const array1: number[] = [];
    const array2: number[] = [];
    
    for await (const line of readLines(file)) {
        array1.push(parseInt(line.split("   ")[0]));
        array2.push(parseInt(line.split("   ")[1]));
    }
    
    array1.sort((a,b) => a - b);
    array2.sort((a,b) => a - b);
    
    let diff: number = 0;
    array1.forEach((value, index) => {
        diff += Math.abs(value - array2[index]);
    });
    
    console.log("Solution1: ", diff);
}


async function solution2() {
    const file = await Deno.open("./input.txt");

    const array1: number[] = [];
    const array2: number[] = [];
    
    for await (const line of readLines(file)) {
        array1.push(parseInt(line.split("   ")[0]));
        array2.push(parseInt(line.split("   ")[1]));
    }
    
    array1.sort((a,b) => a - b);
    array2.sort((a,b) => a - b);

    array1.forEach(el => {
        const result = array2.flatMap((num, index) => (num === el ? index : []));
        console.log(result); 
    });

}

solution1();
solution2();
