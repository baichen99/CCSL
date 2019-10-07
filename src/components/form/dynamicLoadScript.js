let callbacks = [];

const dynamicLoadScript = (src, callback) => {
  const existingScript = document.getElementById(src);
  const cb = callback || function() {};

  if (!existingScript) {
    const script = document.createElement("script");
    script.src = src; // src url for the third-party library being loaded.
    script.id = src;
    document.body.appendChild(script);
    callbacks.push(cb);
    script.onload = function() {
      this.onerror = this.onload = null;
      for (const cb of callbacks) {
        cb(null, script);
      }
      callbacks = null;
    };
    script.onerror = function() {
      this.onerror = this.onload = null;
      cb(new Error("Failed to load " + src), script);
    };
  }

  if (existingScript && cb) {
    if (window.tinymce) {
      cb(null, existingScript);
    } else {
      callbacks.push(cb);
    }
  }
};

export default dynamicLoadScript;
