namespace Frontend.Models;

public record GameLobby(string Name, string System, string[] Players);

public record Friend(string Name, string IdkNow);

public record DashboardData(IReadOnlyList<GameLobby> Games, IReadOnlyList<Friend> Friends);
