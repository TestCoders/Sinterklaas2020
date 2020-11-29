# Sinterklaas2020
Door corona heeft de Sint financieel een zwaar jaar <br>
Echter wil hij niet eindigen als een bedelaar<br>
Hij moet regelmatig zijn kredietlimiet checken<br>
Daarom is het tijd dat er wat Pieten vertrekken<br>


Door de coronamaatregelen heeft ook de Sint een zwaar jaar. Hij heeft wat Inkoop Pieten moeten laten gaan... Om flink wat kosten te besparen, heeft hij Computer Piet de opdracht gegeven om een stukje van de cadeau-inkoop te automatiseren. De opdracht is dat er connectors gebouwd gaan worden. Deze connectors checken bij de leveranciers wie een cadeau het goedkoopst aanbiedt. Vervolgens wordt er bij de goedkoopste aanbieder een verzoek gedaan om een cadeau aan te schaffen. Alle responses die de connectors teruggeven, moeten worden gemockt!<br>
Clone de repo, maak een Pull Request, maar NIET MERGEN


<h2>Opdracht 1</h2>
  <ul>Zet 3 connectors op naar de onderstaande API's</ul>
  <ul>Zet een unittestproject op</ul>

<h2>Opdracht 2</h2>
  <ul>Zorg ervoor dat je connectors een POST en een GET request kunnen uitvoeren. Onderstaand het model hoe de responses van de bedrijven eruit zien</ul>
  <ul>Maak unit tests om de connectors te gebruiken. <b>Mock de response van de POST met een status 200 en mock de response van de GET met het onderstaand model</b></ul>

<h2>Opdracht 3</h2>
  <ul>Maak nu logica die checkt bij welke aanbieder het cadeau het goedkoopst aangeboden wordt. Vervolgens wordt er naar deze aanbieder de POST request verzonden</ul>
  <ul>Toon dit aan met integratietests (wel nog steeds met gemockte resposes)</ul>


<h2>Aanbieders</h2>
Bollie.com    https://www.bollie.com/cadeau/{id}    <br>
CoolBère.com  https://www.coolbere.com/cadeau/{id}  <br>
Aliblabla.com https://www.aliblabla.com/cadeau/{id}  <br>

<h2>Model GET response voor Bollie en CoolBère.</h2>
{
   "product":{
      "Id":5,
      "Price":4.89,
      "Name":"Playdebiel"
   }
}

<h3>Optioneel: Aliblabla geeft nu het bericht terug in XML!</h3>
