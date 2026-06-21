// CWE-79: Second-Order Cross-Site Scripting
// Vulnerability: user-supplied 'displayName' is HTML-encoded at write time
// (looks safe), but decoded back to raw HTML when loaded for the admin template
// renderer which calls Html.Raw(). Fixing the admin renderer alone misses that
// other renderers also call Html.Raw() on the same field in edge cases.

using System;
using System.Web;
using System.Collections.Generic;

public class UserProfileService
{
    private static readonly Dictionary<string, string> _profiles = new();

    // Write path: encodes on save — developer considers this 'sanitized'
    public static void SaveProfile(string userId, string displayName, string bio)
    {
        // Encoding at storage: looks safe but creates false trust downstream
        var encoded = HttpUtility.HtmlEncode(displayName);
        _profiles[userId] = $"{encoded}||{bio}";
    }

    // Read path: decodes back to raw HTML for 'template flexibility'
    public static (string DisplayName, string Bio) LoadProfile(string userId)
    {
        if (!_profiles.TryGetValue(userId, out var raw))
            return (string.Empty, string.Empty);

        var parts = raw.Split("||', 2);
        // VULNERABLE: HtmlDecode reverses the encoding done at save time,
        // restoring attacker-controlled HTML before it reaches the renderer
        return (HttpUtility.HtmlDecode(parts[0]), parts.Length > 1 ? parts[1] : string.Empty);
    }

    // Admin renderer: trusts LoadProfile output as 'already processed'
    public static string RenderAdminCard(string userId)
    {
        var (name, bio) = LoadProfile(userId);
        // Html.Raw equivalent — injects decoded name directly into markup
        return $"<div class='admin-card'><h2>{name}</h2><p>{HttpUtility.HtmlEncode(bio)}</p></div>";
        //                                    ^^^^ VULNERABLE: name is decoded, not re-encoded
    }
}
