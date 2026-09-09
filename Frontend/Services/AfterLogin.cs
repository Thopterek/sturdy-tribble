using Frontend.Models;

namespace Frontend.Services;

/*
 * Bunch of hardcoded fake elements to be replaced
 * to use for the actual call to the program
*/
public class AfterLogin
{
    public static readonly Guid SkavenId = new("11111111-1111-1111-1111-111111111111");
    public static readonly Guid CtululuId = new("21111111-1111-1111-1111-111111111111");
    public IReadOnlyList<GameLobby> all_games =
    [
        new GameLobby(
            SkavenId,
            "Skavens going to war",
            "Warhammer 5ed",
            ["Dawcio", "Wojtas", "Kuba"]
        ),
        new GameLobby(
            CtululuId,
            "Investigation of cosmic powers",
            "Homebrew Cthullu",
            ["Matloszek", "Olek"]
        ),
    ];

    public async Task<GameLobby?> GetLobby(Guid? lobby_id)
    {
        if (lobby_id is null)
            return null;
        foreach (GameLobby gl in all_games)
        {
            if (gl.Id == lobby_id)
                return gl;
        }
        return null;
    }

    public async Task<DashboardData> GetDashboardDataAsync()
    {
        return new DashboardData(
            Games: all_games,
            Friends: [new Friend("Dawcio", true, 53), new Friend("Matloszek", false, 2532)],
            new Friend("Stelio", true, 4)
        );
    }
}
