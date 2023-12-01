const std = @import("std");
const print = @import("std").debug.print;


const digits = [_][]const u8{ "one", "two", "three", "four", "five", "six", "seven", "eight", "nine" };

fn solution(line: []const u8) u8 {
    var digitOne: ?u8 = null;
    var digitTwo: ?u8 = null;
    for (line, 0..) |c, i| {
        var digit: ?u8 = null;
        if (c >= '0' and c <= '9') {
            digit = c - '0';
        } else {
            for (digits, 1..) |numStr, d| {
                const lineSlice = line[i..@min(i + numStr.len, line.len)];
                if (std.mem.eql(u8, numStr, lineSlice)) {
                    digit = @intCast(d);
                }
            }
        }

        if (digit) |d| {
            _ = d;
            if (digitOne == null) {
                digitOne = digit;
            }
            digitTwo = digit;
        }
    }
    return 10 * digitOne.? + digitTwo.?;
}

pub fn main() !void {
    
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    var sum: u64 = 0;

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        sum += solution(line);
    }   

    print("{}", .{sum});

}