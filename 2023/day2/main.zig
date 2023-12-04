const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();
    
    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    
    var buf: [1024]u8 = undefined;
    const red_balls: u8 = 12;
    _ = red_balls;
    const green_balls: u8 = 13;
    _ = green_balls;
    const blue_balls: u8 = 14;
    _ = blue_balls;
    
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        print("{s}\n", .{line});
    
        //only 12 red cubes, 13 green cubes, and 14 blue cubes
        var dotdot_split = std.mem.split(u8, line, ":");
        while (dotdot_split.next()) |split| {
            print("{s}\n", .{split});
        }
    }

}