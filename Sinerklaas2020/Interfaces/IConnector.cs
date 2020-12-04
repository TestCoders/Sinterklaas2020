using Sinerklaas2020.Models;
using System.Net;
using System.Threading.Tasks;

namespace Sinerklaas2020.Interfaces
{
    public interface IConnector
    {
        Task<Product> Get(int id, string url);
        Task<HttpStatusCode> Post(int id, string url);
    }
}
