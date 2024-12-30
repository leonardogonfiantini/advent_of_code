import {readLines} from "https://deno.land/std/io/mod.ts";

async function parseInput(filePath: string): Promise<number[][]> {

  const file = await Deno.open(filePath);

  const matrix: number[][] = [];
  for await (const line of readLines(file)) {
    const parsedLine: number[] = line.split(" ").map(function (val) { 
                                      return parseInt(val) 
                                  });
    matrix.push(parsedLine);
  }

  return matrix;
}

const input: number[][] = await parseInput('input.txt');


function checkReport(input: number[]): number {
  
  let descORasc: number = 0;
  if (input[0] > input[1]) descORasc = 1;

  for (let i = 0; i < input.length - 1; i++) {
    if (descORasc == 0 && input[i] > input[i+1]) return i;
    if (descORasc == 1 && input[i] < input[i+1]) return i;

    const diff: number = Math.abs(input[i] - input[i+1])
    if (diff < 1 || diff > 3) return i;
  }

  return -1;
}

function solution1(input: number[][]): number {

  let score: number = 0;

  for (const report of input) {
    if (checkReport(report) != -1) continue;
    score +=1;
  }

  return score;
}

function solution2(input: number[][]): number {
  
  
  let score: number = 0;

  for (const report of input) {
    for (let i = 0; i < report.length; i++) {
      const temp_report = report.filter((_, index) => index !== i)
      if (checkReport(temp_report) === -1) {
        score += 1 
        break;
      }
    }
  }

  return score;

}


console.log("Solution1: ", solution1(input));
console.log("Solution2: ", solution2(input));