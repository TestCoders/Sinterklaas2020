using Moq;
using NUnit.Framework;
using Sinerklaas2020.Connectors;
using Sinerklaas2020.Interfaces;
using Sinerklaas2020.Models;
using Sinerklaas2020.Services;
using System.Threading.Tasks;

namespace Sinterklaas2020.Tests.IntegrationTests
{
    public class IntegrationTests : Init
    {
        private Mock<IConnector> _mockConnector;
        private Mock<CompanyConnector> _mockBollieConnector;
        private Mock<CompanyConnector> _mockCoolBereConnector;
        private Mock<CompanyConnector> _mockAliBlaBlaConnector;

        [SetUp]
        public void SetUpIntegrationTests()
        {
            // Instantiate mocks
            _mockConnector = new Mock<IConnector>();
            _mockBollieConnector = new Mock<CompanyConnector>(_mockConnector.Object, "https://www.bollie.com/cadeau");
            _mockCoolBereConnector = new Mock<CompanyConnector>(_mockConnector.Object, "https://www.coolbere.com/cadeau");
            _mockAliBlaBlaConnector = new Mock<CompanyConnector>(_mockConnector.Object, "https://www.aliblabla.com/cadeau");

            // Setup returned products
            _mockBollieConnector.Setup(mock => mock.GetProduct(5)).Returns(Task.FromResult(new Product { Id = 5, Price = 4.99, Name = "Playdebiel" }));
            _mockCoolBereConnector.Setup(mock => mock.GetProduct(5)).Returns(Task.FromResult(new Product { Id = 5, Price = 5.11, Name = "Playdebiel" }));
            _mockAliBlaBlaConnector.Setup(mock => mock.GetProduct(5)).Returns(Task.FromResult(new Product { Id = 5, Price = 3.99, Name = "Playdebiel" }));
        }

        [Test]
        public async Task GivenThreeProducts_GetCheapestConnector_ShouldBeLowestPriceConnector()
        {
            // Assemble
            var giftService = new BuyGiftService(_mockBollieConnector.Object, _mockCoolBereConnector.Object, _mockAliBlaBlaConnector.Object );

            // Act
            var result = await giftService.GetCheapestConnector();

            // Assert
            Assert.AreEqual("https://www.aliblabla.com/cadeau", result.Url);
        }

        [Test]
        public async Task GivenThreeProducts_BuyProductWithCheapestConnector_ShouldBeLowestPriceConnector()
        {
            // Assemble
            var giftService = new BuyGiftService(_mockBollieConnector.Object, _mockCoolBereConnector.Object, _mockAliBlaBlaConnector.Object);
            
            // Assign
            var result = await giftService.BuyCheapestGift(5);

            // Act
            Assert.That(result.ToLower().Contains("aliblabla"));
        }
    }
}
