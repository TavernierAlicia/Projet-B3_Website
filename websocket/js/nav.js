function createNav() {
  document.write(`
    <!DOCTYPE html>
    <html lang="fr">
      <head>
        <meta charset="utf-8">
        <link rel="icon" type="image/png" href="../pictures/favicon.png" />
        <title>Order'N Drink - Accueil</title>
      </head>
      <body>
        <nav>
          <a class="index" href="/index"><img src="../pictures/logo.png" alt="logo Order'N Drink"></a>
          <div class="navbar-right-items">
	    <button id="menuButton">Menu</button>
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

