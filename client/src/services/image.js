import Pica from "pica";

const dataURLToImage = async bytes => {
  return await new Promise(resolve => {
    var img = new Image();
    img.src = bytes;
    img.onload = async () => {
      resolve(img);
    };
  });
};

const blobToDataURL = blob => {
  return new Promise(resolve => {
    var a = new FileReader();
    a.onload = function(e) {
      resolve(e.target.result);
    };
    a.readAsDataURL(blob);
  });
};

export const getSize = async dataUrl => {
  const image = await dataURLToImage(dataUrl);
  return {
    width: image.width,
    height: image.height
  };
};

export const toBlob = dataUrl => {
  var arr = dataUrl.split(","),
    mime = arr[0].match(/:(.*?);/)[1],
    bstr = atob(arr[1]),
    n = bstr.length,
    u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new Blob([u8arr], { type: mime });
};

export const resize = async (blob, opts) => {
  const dataUrl = await blobToDataURL(blob);
  const original = await dataURLToImage(dataUrl);
  const resizedCanvas = document.createElement("canvas");
  resizedCanvas.height = opts.height;
  resizedCanvas.width = opts.width;
  const p = Pica();
  try {
    const result = await p.resize(original, resizedCanvas, {
      unsharpAmount: 80,
      unsharpRadius: 0.6,
      unsharpThreshold: 2
    });
    return toBlob(result.toDataURL());
  } catch (err) {
    // TODO: Add logging
    return;
  }
};
