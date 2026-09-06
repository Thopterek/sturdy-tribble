namespace Frontend.Models;

public record GameLobby(string Name, string System, string[] Players);

public record Friend(string Name, bool Online, int Hours);

public record DashboardData(
    IReadOnlyList<GameLobby> Games,
    IReadOnlyList<Friend> Friends,
    Friend User
);
