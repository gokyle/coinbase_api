/*
   Package coinbase_api is a Go interface to the CoinBase.com
   API. It may be used to design automated Bitcoin trading
   systems.

   Currently implemented endpoints:
        * get account balance
        * get receive address
        * get exchange rate
        * purchase bitcoins
        * sell bitcoins

   The global variable `ApiKey` is used to store your Coinbase API
   key. If this is empty, attempts to make authenticated requests
   will result in an ErrNotAuthenticated error being returned.
*/
/*
   Copyright (c) 2013 Kyle Isom <kyle@tyrfingr.is>

   Permission to use, copy, modify, and distribute this software for any
   purpose with or without fee is hereby granted, provided that the above
   copyright notice and this permission notice appear in all copies.

   THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
   WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
   MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
   ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
   WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
   ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
   OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
*/
package coinbase_api
