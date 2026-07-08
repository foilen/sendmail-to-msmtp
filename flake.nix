{
  description = "sendmail-to-msmtp: a bridge from sendmail to msmtp";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "sendmail-to-msmtp";
          version = "1.1.1";

          src = ./.;

          subPackages = [ "sendmail" ];

          vendorHash = null;

          postInstall = ''
            mv $out/bin/sendmail $out/bin/sendmail-to-msmtp
          '';

          meta = with pkgs.lib; {
            description = "A bridge from sendmail to msmtp";
            homepage = "https://github.com/foilen/sendmail-to-msmtp";
            license = licenses.mit;
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = [ pkgs.go ];
        };
      });
}
