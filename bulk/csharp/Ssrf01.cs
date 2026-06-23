// CWE-918: SSRF via HttpClient with user-controlled URL
using System.Net.Http;
public class Ssrf01 {
    private static readonly HttpClient _client = new();
    public async Task<string> Fetch(string url) {
        return await _client.GetStringAsync(url); // no allowlist
    }
}
