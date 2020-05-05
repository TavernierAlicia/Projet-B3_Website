function createNav() {
  document.write(`
    <!DOCTYPE html>
    <html lang="fr">
      <head>
        <meta charset="utf-8">
        <title>Order'N Drink - Accueil</title>
      </head>
      <body>
        <nav>
          <a class="index" href="/index"><img src="../pictures/logo.png"></a>
          <div class="navbar-right-items">
            <ul>
              <li class="nav-item">
                <a class="nav-link" href="/professionnal">Devenir partenaire</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/form">Contact</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/about">À propos</a>
              </li>
            </ul>
          </div>
        </nav>
      </body>
    </html>
  `);
}

createNav();

