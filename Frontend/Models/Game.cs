namespace Frontend.Models;

public record LobbyMessage(string SetPlayer, string SetMsg)
{
    public readonly string player = SetPlayer;
    public readonly string msg = SetMsg;
    public readonly DateTime send_at = DateTime.Now;
}
