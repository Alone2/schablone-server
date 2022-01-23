from typing import Any, Dict, Optional

import httpx

from ...client import AuthenticatedClient
from ...models.macro import Macro
from ...types import Response


def _get_kwargs(
    macro_id: int,
    *,
    client: AuthenticatedClient,
) -> Dict[str, Any]:
    url = "{}/macro/get/{macroId}".format(client.base_url, macroId=macro_id)

    headers: Dict[str, str] = client.get_headers()
    cookies: Dict[str, Any] = client.get_cookies()

    return {
        "method": "get",
        "url": url,
        "headers": headers,
        "cookies": cookies,
        "timeout": client.get_timeout(),
    }


def _parse_response(*, response: httpx.Response) -> Optional[Macro]:
    if response.status_code == 200:
        response_200 = Macro.from_dict(response.json())

        return response_200
    return None


def _build_response(*, response: httpx.Response) -> Response[Macro]:
    return Response(
        status_code=response.status_code,
        content=response.content,
        headers=response.headers,
        parsed=_parse_response(response=response),
    )


def sync_detailed(
    macro_id: int,
    *,
    client: AuthenticatedClient,
) -> Response[Macro]:
    """Get a macro

    Args:
        macro_id (int):

    Returns:
        Response[Macro]
    """

    kwargs = _get_kwargs(
        macro_id=macro_id,
        client=client,
    )

    response = httpx.request(
        verify=client.verify_ssl,
        **kwargs,
    )

    return _build_response(response=response)


def sync(
    macro_id: int,
    *,
    client: AuthenticatedClient,
) -> Optional[Macro]:
    """Get a macro

    Args:
        macro_id (int):

    Returns:
        Response[Macro]
    """

    return sync_detailed(
        macro_id=macro_id,
        client=client,
    ).parsed


async def asyncio_detailed(
    macro_id: int,
    *,
    client: AuthenticatedClient,
) -> Response[Macro]:
    """Get a macro

    Args:
        macro_id (int):

    Returns:
        Response[Macro]
    """

    kwargs = _get_kwargs(
        macro_id=macro_id,
        client=client,
    )

    async with httpx.AsyncClient(verify=client.verify_ssl) as _client:
        response = await _client.request(**kwargs)

    return _build_response(response=response)


async def asyncio(
    macro_id: int,
    *,
    client: AuthenticatedClient,
) -> Optional[Macro]:
    """Get a macro

    Args:
        macro_id (int):

    Returns:
        Response[Macro]
    """

    return (
        await asyncio_detailed(
            macro_id=macro_id,
            client=client,
        )
    ).parsed
