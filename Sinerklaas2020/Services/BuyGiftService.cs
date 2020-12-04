using Sinerklaas2020.Connectors;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Sinerklaas2020.Services
{
    public class BuyGiftService
    {
        private readonly CompanyConnector[] _connectorArray;

        public BuyGiftService(params CompanyConnector[] companyConnectors)
        {
            _connectorArray = companyConnectors;
        }

        public async Task<CompanyConnector> GetCheapestConnector()
        {
            var connectorPriceCollection = new Dictionary<CompanyConnector, double>();

            foreach (var connector in _connectorArray)
            {
                var result = await connector.GetProduct(5);
                var price = result.Price;
                connectorPriceCollection.Add(connector, price);
            }

            var chepeastPrice = connectorPriceCollection.Values.Min();
            return connectorPriceCollection.FirstOrDefault(product => product.Value == chepeastPrice).Key;
        }

        public async Task<string> BuyCheapestGift(int id)
        {
            var connector = await GetCheapestConnector();
            await connector.BuyProduct(id);
            return $"Product bought at: {connector.Url}";
        }
    }
}
