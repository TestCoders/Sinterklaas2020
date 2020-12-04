using Moq;
using NUnit.Framework;
using Sinerklaas2020.Connectors;
using Sinerklaas2020.Interfaces;
using Sinerklaas2020.Models;
using System.Net;
using System.Threading.Tasks;

namespace Sinterklaas2020.Tests.UnitTests
{
    public class Connector : Init
    {
        private readonly string _urlBollie = "https://www.bollie.com/cadeau";
        private Product _product;
        private Mock<IConnector> _mockConnector;

        [SetUp]
        public void SetUpUnitTests()
        {
            // Assign product
            _product = new Product
            {
                Id = 5,
                Price = 4.99,
                Name = "Playdebiel"
            };

            // Assemble mock
            _mockConnector = new Mock<IConnector>();
            _mockConnector.Setup(mock => mock.Get(_product.Id, _urlBollie)).Returns(Task.FromResult(_product));
            _mockConnector.Setup(mock => mock.Post(_product.Id, _urlBollie)).Returns(Task.FromResult(HttpStatusCode.OK));
        }

        [Test]
        public async Task BollieGet_CorrectId_ShouldReturnProduct()
        {
            // Assemble
            var bollieConnector = new CompanyConnector(_mockConnector.Object, _urlBollie);

            // Act
            var result = await bollieConnector.GetProduct(_product.Id);

            // Assert
            Assert.AreEqual(_product.Price, result.Price);
            _mockConnector.Verify(mock => mock.Get(_product.Id, _urlBollie), Times.Once);
            _mockConnector.Verify(mock => mock.Post(_product.Id, _urlBollie), Times.Never);
        }

        [Test]
        public async Task BolliePut_CorrectId_ShouldReturnStatusCode200()
        {
            // Assemble
            var bollieConnector = new CompanyConnector(_mockConnector.Object, _urlBollie);

            // Act
            var result = await bollieConnector.BuyProduct(5);

            // Assert
            Assert.AreEqual(HttpStatusCode.OK, result);
            _mockConnector.Verify(mock => mock.Get(_product.Id, _urlBollie), Times.Never);
            _mockConnector.Verify(mock => mock.Post(_product.Id, _urlBollie), Times.Once);
        }
    }
}
