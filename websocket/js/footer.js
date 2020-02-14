function createFooter() {
  document.write(`
  <!DOCTYPE html>
  <html lang="fr">
    <head>
      <meta charset="utf-8">
      <title>Order'N Drink - Accueil</title>
      <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
    </head>
    <body>
    </body>
    <footer>
      <div class="container-fluid">
        <div class="container">
          <div class="row">
            <div class="col">
              <a href="/about">À propos</a>
            </div>
            <div class="col">
              <a href="/contact"> Contact </a>
            </div>
            <div class="col">
              <a href="/about"> Mentions légales </a>
            </div>
          </div>
        </div>
        <p style="text-align: center;">Order'N Drink - Copyright 2019</p>
      </div>
    </footer>
  </html>
  `);

}

createFooter();