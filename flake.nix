{
  description = "Pollen dispersion data CLI from WeatherNews API";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = nixpkgs.legacyPackages.${system};
        in
        {
          default = pkgs.buildGoModule {
            pname = "pollenso";
            version = "0.1.0";
            src = ./.;
            vendorHash = "sha256-pc0iLRslEwnKbRYrQJBMD1w0flr7ZweJ8iBDvSk/w2M=";
            meta = {
              description = "Pollen dispersion data CLI from WeatherNews API";
              license = pkgs.lib.licenses.mit;
              mainProgram = "pollenso";
            };
          };
        }
      );
    };
}
