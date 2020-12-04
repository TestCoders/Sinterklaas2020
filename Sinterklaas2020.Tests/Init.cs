using Microsoft.Extensions.DependencyInjection;
using NUnit.Framework;
using Sinerklaas2020.Connectors;
using Sinerklaas2020.Interfaces;

namespace Sinterklaas2020.Tests
{
    public class Init
    {
        public IConnector Connector;

        [SetUp]
        public void SetUp()
        {
            var serviceCollection = new ServiceCollection();
            serviceCollection.AddScoped<IConnector, Connector>();
            var serviceProvider = serviceCollection.BuildServiceProvider();

            Connector = serviceProvider.GetService<IConnector>();
        }
    }
}
