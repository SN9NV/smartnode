{
  description = "The CLI package for Rocket Pool smart nodes";

  inputs.nixpkgs.url = "nixpkgs/nixos-22.11";

  outputs = { self, nixpkgs }:
  let
    supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
    forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
    nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
  in
  {

    packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
    rocketpool-cli = pkgs.buildGoModule {
      pname = "rocketpool";
      CGO_ENABLED = false;
      version = "v1.7.5";
      src = ./.;
      vendorSha256 = "sha256-RKxPXSGhqednUVZQSXyzRRx4aI2hmmqANaeuO5tQiF0=";
      subPackages = ["rocketpool-cli"];
    };
});

    defaultPackage = forAllSystems (system: self.packages.${system}.rocketpool-cli);


  };
}
