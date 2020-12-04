using Newtonsoft.Json;
using Sinerklaas2020.Interfaces;
using Sinerklaas2020.Models;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace Sinerklaas2020.Connectors
{
    public class Connector : IConnector
    {
        private readonly HttpClient _client = new HttpClient();

        public async Task<Product> Get(int id, string url)
        {
            var result = await _client.GetAsync($"{url}/{id}");
            return JsonConvert.DeserializeObject<Product>(result.Content.ReadAsStringAsync().Result);
        }

        public async Task<HttpStatusCode> Post(int id, string url)
        {
            var result = await _client.PostAsync($"{url}/{id}", new StringContent($"{id}", Encoding.UTF8, "application/json"));
            return result.StatusCode;
        }
    }
}
