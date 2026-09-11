namespace Frontend.Models;

public class Game
{
    public DateTime last_update;
}

public record LobbyMessage(string SetPlayer, string SetMsg)
{
    public readonly string player = SetPlayer;
    public readonly string msg = SetMsg;
    public readonly DateTime send_at = DateTime.Now;
}
