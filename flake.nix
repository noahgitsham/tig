{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }: let 
    pkgs = nixpkgs.legacyPackages.x86_64-linux;
    common-deps = with pkgs; [ libcap go gcc ];
  in {

    devShell.x86_64-linux = pkgs.mkShell {
      buildInputs = with pkgs; [  ] ++ common-deps;
    };

    packages.x86_64-linux.tig-cli = pkgs.buildGoPackage {
      goDeps = with pkgs; [  ] ++ common-deps;
    };

    packages.x86_64-linux.tig-server = self.packages.x86_64-linux.hello;
  };
}
