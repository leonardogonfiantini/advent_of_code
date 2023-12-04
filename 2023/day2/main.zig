const std = @import("std");
const print = std.debug.print;

pub fn sol1() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    const red_balls: u64 = 12;
    const green_balls: u64 = 13;
    const blue_balls: u64 = 14;

    var sum: u64 = 0;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var flag: u64 = 0;
        //only 12 red cubes, 13 green cubes, and 14 blue cubes
        var dotdot_split = std.mem.split(u8, line, ":");
        const split_space = dotdot_split.next().?;
        var game_name = std.mem.split(u8, split_space, " ");
        _ = game_name.next();
        const game_value = game_name.next().?;

        while (dotdot_split.next()) |dotdot| {
            var dotcomma_split = std.mem.split(u8, dotdot, ";");
            while (dotcomma_split.next()) |dotcomma| {
                var comma_split = std.mem.split(u8, dotcomma, ",");
                while (comma_split.next()) |comma| {
                    var space_split = std.mem.split(u8, comma, " ");
                    _ = space_split.next();
                    const ball_value: u64 = try std.fmt.parseInt(u64, space_split.next().?, 10);
                    const color = space_split.next().?;

                    if (std.mem.eql(u8, color, "red") and ball_value > red_balls) {
                        flag = 1;
                    }

                    if (std.mem.eql(u8, color, "green") and ball_value > green_balls) {
                        flag = 1;
                    }

                    if (std.mem.eql(u8, color, "blue") and ball_value > blue_balls) {
                        flag = 1;
                    }
                }
            }
        }

        if (flag == 0) {
            const val = try std.fmt.parseInt(u64, game_value, 10);
            sum += val;
        }
    }

    print("{}", .{sum});
}


pub fn sol2() !void {
    print("Hello world\n", .{});
}

pub fn main() !void {
    try sol1();
}
