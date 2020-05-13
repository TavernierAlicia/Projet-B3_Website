function createFooter() {
  document.write(`
  <!DOCTYPE html>
  <html lang="fr">
    <head>
      <meta charset="utf-8">
      <title>Order'N Drink - Accueil</title>
    </head>
    <body>
    </body>
    <footer>
      <div class="container">
      	<div class="wrapper">
        <div class="footer-left">
          <div class="footer-line">
            <a href="/about">À propos</a>
          </div>
          <div class="footer-line">
            <a href="/form"> Contact </a>
          </div>
          <div class="footer-line">
            <a href="/professionnal"> Devenir partenaire </a>
          </div>
          <div class="footer-line">
            <a href="/legal"> Mentions légales </a>
          </div>
        </div>

        <div class="footer-middle">
          <div class="footer-line">
            <p>Une question?</p>
            <p>service.orderndrink@gmail.com</p>
          </div>
          <div class="footer-line">
            <p> 01 76 09 34 08 </p>
          </div>
          <div class="footer-line">
            <a href="#"><img class=social src="../pictures/fb.png"></a>
            <a href="#"><img class=social src="../pictures/insta.png"></a>
          </div>
        </div>

        <div class="footer-right">
          <div class="footer-line">
            <a class="footer-store" href="#"><img class="white-store" src="../pictures/app-store-blanc.png"></a>
          </div>
          <div class="footer-line">
            <a class="footer-store" href="#"><img class="white-store" src="../pictures/google-play-blanc.png"></a>
          </div>
        </div>
      </div>
      </div>
      <div class="cp">
        <p>Order'N Drink - Copyright 2019</p>
      </div>
    </footer>
  </html>
  `);

}

createFooter();
