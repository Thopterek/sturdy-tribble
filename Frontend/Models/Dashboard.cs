namespace Frontend.Models;

public record GameLobby(Guid Id, string Name, string System, List<Friend> Players, int Hours);

public record Friend(string Name, bool Online, int Hours);

public record DashboardData(
    IReadOnlyList<GameLobby> Games,
    IReadOnlyList<Friend> Friends,
    Friend User
);
