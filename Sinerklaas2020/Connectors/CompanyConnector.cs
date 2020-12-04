using Sinerklaas2020.Interfaces;
using Sinerklaas2020.Models;
using System.Net;
using System.Threading.Tasks;

namespace Sinerklaas2020.Connectors
{
    public class CompanyConnector
    {
        public readonly string Url;
        private readonly IConnector _connector;

        public CompanyConnector(IConnector connector, string url)
        {
            _connector = connector;
            Url = url;
        }

        public virtual async Task<Product> GetProduct(int id)
        {
            return await _connector.Get(id, Url);
        }

        public async Task<HttpStatusCode> BuyProduct(int id)
        {
            var result = await _connector.Post(id, Url);
            return result;
        }
    }
}
