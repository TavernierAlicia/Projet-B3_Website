'use strict';

//create regexes
const nameReg = new RegExp('^[A-Za-zÀ-ÿ\s -]+$');
const mailReg = new RegExp('^([a-zA-ZÀ-ÿ0-9_\.\-])+\@(([a-zA-ZÀ-ÿ0-9\-])+\\.)+([a-zA-ZÀ-ÿ0-9]{2,4})+$');
const telReg = new RegExp('^([0-9]{2}){4}[0-9]{2}');

window.onload = init;
function init() {

  //get and disable button
  const button = document.getElementById('sendform');
  button.disabled = true;
}


function checkName() {
  var button = document.getElementById('sendform');
  var name = document.getElementById("name").value;
  var pname = document.getElementById("errname");

  if (name != '') {
    if (nameReg.exec(name)) {
      pname.innerText = '';
      return true;
    } else {
      pname.innerText = 'Nom invalide';
      button.disabled = true;
      document.getElementById('reglement').checked = false;
      return false;
    }
  } else {
    pname.innerText = 'Champ \"Nom\" vide';
    button.disabled = true;
    document.getElementById('reglement').checked = false;
    return false;
  }
}

function checkSurname() {
  var button = document.getElementById('sendform');
  var surname = document.getElementById("surname").value;
  var psurname = document.getElementById("errsurname");

  if (surname != '') {
    if (nameReg.exec(surname)) {
      psurname.innerText = '';
      return true
    } else {
      psurname.innerText = 'Prenom invalide';
      button.disabled = true;
      document.getElementById('reglement').checked = false;
      return false;
    }
  } else {
    psurname.innerText = 'Champ \"Prenom\" vide';
    button.disabled = true;
    document.getElementById('reglement').checked = false;
    return false;
  }
}

function checkMail() {
  var button = document.getElementById('sendform');
  var mail = document.getElementById("email").value;
  var pmail = document.getElementById("erremail");

  if (mail != '') {
    if (mailReg.exec(mail)) {
      pmail.innerText = '';
      return true;
    } else {
      pmail.innerText = 'Addresse mail invalide';
      button.disabled = true;
      document.getElementById('reglement').checked = false;
      return false;
    }
  } else {
    pmail.innerText = 'Champ \"Mail\" vide';
    button.disabled = true;
    document.getElementById('reglement').checked = false;
    return false;
  }
}

function checkNumber() {
  var button = document.getElementById('sendform');
  var tel = document.getElementById("tel").value;
  var ptel = document.getElementById("errtel");

  if (tel != '') {
    if (telReg.exec(tel)) {
      ptel.innerText = '';
      return true;
    } else {
      ptel.innerText = 'Numero de telephone invalide';
      button.disabled = true;
      document.getElementById('reglement').checked = false;
      return false;
    }
  } else {
    ptel.innerText = 'Champ \"Numero de telephone\" vide';
    button.disabled = true;
    document.getElementById('reglement').checked = false;
    return false;
  }
}

function checkMessage() {
  var button = document.getElementById('sendform');
  var message = document.getElementById("message").value;
  var pmessage = document.getElementById("errmessage");

  if (message != '') {
    if (message.length > 20 && message.length <= 500) {
      pmessage.innerText = '';
      return true;
    } else if (message.length > 500) {
      pmessage.innerText = 'Message invalide (max 500 characteres)';
      button.disabled = true;
      document.getElementById('reglement').checked = false;
      return false;
    } else {
      pmessage.innerText = 'Message invalide (min 20 characteres)';
      button.disabled = true;
      document.getElementById('reglement').checked = false;
      return false;
    }
  } else {
    pmessage.innerText = 'Champ \"Message\" vide';
    button.disabled = true;
    document.getElementById('reglement').checked = false;
    return false;
  }
}

function checkAll() {
  var button = document.getElementById('sendform');
  if (checkMail() && checkName() && checkSurname() && checkMessage() && checkNumber() && validate()) {
    button.disabled = false;
  } else {
    document.getElementById('reglement').checked = false;
    button.disabled = true;
  }
}

function validate() {
  var checkbox = document.getElementById('reglement').checked;
  if (checkbox == true) {
      return true;
  } else {
    return false;
  }
}

