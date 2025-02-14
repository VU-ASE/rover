// The host of the server
// const serverHost = "localhost";
const serverHost = "localhost";
// The port of the server
const serverPort = 7500;
// If we enable https
const serverScheme = "http";
// For easy access
const serverAddress = `${serverScheme}://${serverHost}:${serverPort}`;

// The ICE servers to use (comment out when on LAN)
const iceServers: {
  urls: string[];
}[] = [
  // {
  //   urls: [
  //     "stun:stun.l.google.com:19302",
  //     "stun:stun1.l.google.com:19302",
  //     "stun:stun2.l.google.com:19302",
  //     "stun:stun3.l.google.com:19302",
  //     "stun:stun4.l.google.com:19302",
  //     "stun:stun.ekiga.net",
  //     "stun:stun.ideasip.com",
  //     "stun:stun.rixtelecom.se",
  //     "stun:stun.schlund.de",
  //   ],
  // },
];

export { serverHost, serverPort, serverScheme, serverAddress, iceServers };
