<script>
    import BgElementParent from "$lib/components/bg-element-parent.svelte";
    import Footer from "$lib/components/Footer.svelte";
    import Nav from "$lib/components/Nav.svelte";
    import { CloudUpload, ClipboardPenLine, Paperclip, FileType2 } from '@lucide/svelte';
    import "../app.css";
  
    let fileInfo = null;
  let file = null; // <-- Tambahkan ini

  function handleFile(event) {
    file = event.target.files[0]; // <-- Simpan file-nya di variabel global
    if (file) {
      fileInfo = {
        name: file.name,
        size: file.size,
        type: file.type
      };
    }
  }

  async function uploadFile() {
    const formData = new FormData();
    formData.append("nmfile", file); // <-- sekarang file tidak undefined

    const res = await fetch(" http://127.0.0.1:3000/upload", {
      method: "POST",
      body: formData
    });

    const blob = await res.blob();
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.href = url;
    a.download = `cnv-${file.name}`;
    a.click();
    URL.revokeObjectURL(url);
  }
  </script>
  
  <Nav />
  
  <div class="w-full h-[100vh] pt-16 flex justify-center items-center bg-black">
    <div class="w-full h-full backdrop-blur-[150px] absolute top-0 left-0 flex flex-col">
      <h1 class="text-7xl text-white mt-16 p-3 text-center">Free Conversi Online</h1>
  
      <div class="w-[450px] h-[400px] border-2 border-dashed border-white m-auto rounded-[50px] bg-black flex flex-col justify-center items-center text-white">
  
        {#if fileInfo}
          <div class="w-full h-auto p-10 flex flex-col">
            <div class="flex">
              <ClipboardPenLine /> <div class="mx-2">:</div>
              <input class="px-2 py-1 border outline-none rounded-md" type="text" bind:value={fileInfo.name} />
            </div>
            <div class="flex mt-3">
              <Paperclip /> <div class="mx-2">:</div>
              {fileInfo.size} bytes
            </div>
            <div class="flex mt-3">
              <FileType2 /> <div class="mx-2">:</div>
              <select class="select border bg-black w-auto">
                <option disabled selected>{fileInfo.type}</option>
                <option>png</option>
                <option>jpeg</option>
                <option>jpg</option>
              </select>
            </div>
  
            <div class="w-full h-auto px-5 flex justify-evenly items-center mt-4">
              <div class="w-[30vw] border"></div>
              <h1 class="bg-black p-3">OR</h1>
              <div class="w-[30vw] border"></div>
            </div>
  
            <div class="w-full flex justify-center items-center">
              <button on:click={uploadFile} class="w-[300px] h-[50px] bg-green-300 rounded-xl flex mt-10">
                <h1 class="m-auto font-bold text-black">CONVERT</h1>
              </button>
            </div>
          </div>
        {:else}
          <div>
            <input type="file" id="usrfile" class="hidden" on:change={handleFile} />
            <label for="usrfile" class="bg-gray-900 py-2 px-4 rounded-md flex justify-between w-36 btn">
              <CloudUpload /> Choose file
            </label>
          </div>
        {/if}
  
      </div>
    </div>
    <BgElementParent />
  </div>
  
  <Footer />
  